package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"douyin/global"
	"douyin/pkg/util"
	"time"

	"github.com/gin-gonic/gin"
)

// Auth 鉴权接口
// 所有接口使用acc token(auth 2h)
// 在Redis中 acc token(2h)为key，ref token(30d)为value
//
// auth 核心过程：
// 1.检查这个acc token，若没过期取uid
// (已解决小问题：res token 可能过期；比如acc token才生成，res token 还剩1h 过期——这种情况不会出现，acc 与 ref 同时更新)
//
// 2.若acc token过期了，取出ref token
// 检查 ref token是否过期，过期叫重新登陆；
// ref token没过期，生成新的acc token，ref token(防止恰好此时失效，同时更新)，删除旧的记录，新建Redis记录
//
// 3.acc和ref同时更新，只有在30天没有没有登录时提醒重新登陆
//

type BackendLoginReq struct {
}

func Auth() gin.HandlerFunc {
	//先判断请求头是否为空，为空则为游客状态
	return func(c *gin.Context) {
		token := ""
		token = c.DefaultQuery("token", "")
		if token == "" {
			token = c.PostForm("token")
		}
		if token == "" {
			c.Set("userId", "")
			c.Next()
			return
		}

		//有token，判断是否过期: 2h
		timeOut, err := util.ValidToken(token)

		if err != nil || timeOut {
			//token过期或者解析token发生错误
			global.Logger.Info(context.Background(), "token expire or parse token error")
			global.Logger.Debug(context.Background(), "valid token err")
			global.Logger.Info(context.Background(), "valid refreshToken")

			// 30d token 能否取出
			value := global.Redis.Get(context.Background(), token)
			refreshToken, err := value.Result()
			if err != nil {
				// debug
				global.Logger.Debug(context.Background(), "get refreshToken from redis err")

				global.Logger.Error(context.Background(), "token不合法，请确认你是否登录")
				c.JSON(200, gin.H{
					"status_code": 400,
					"msg":         "token不合法，请确认你操作是否有误",
				})
				c.Abort()
				return
			}

			// 可以取出30d token, 检查是否过期
			timeOut, err := util.ValidToken(refreshToken)
			if err != nil || timeOut {
				global.Logger.Debug(context.Background(), "valid refreshToken err:")
				//refreshToken出问题，表明用户三十天未登录，需要重新登录
				global.Logger.Info(context.Background(), "user need login again")
				global.Redis.Del(context.Background(), token)
				c.Set("userId", "")
				c.Next()
				return
			}

			// refresh token 没过期
			userId, err := util.GetUserIDFormToken(refreshToken)
			if err != nil {
				global.Logger.Error(context.Background(), "parse token to get uid error:")
				//token解析不了的情况一般很少,暂时panic一下
				panic(err)
			}

			//根据refreshToken 更新 accessToken
			accessToken, err := util.CreateAccessToken(userId)
			if err != nil {
				global.Logger.Error(context.Background(), "create acc token error:")
				//token解析不了的情况一般很少,暂时panic一下
				panic(err)
			}

			//更新后，重新设置redis的key
			newRefreshToken, err := util.CreateRefreshToken(userId)
			if err != nil {
				global.Logger.Error(context.Background(), "creat ref token error:")
				panic(err)
			}

			if err := global.Redis.Set(context.Background(), token, newRefreshToken, 30*24*time.Hour).Err(); err != nil {
				global.Logger.Error(context.Background(), "create redis acc token error")
			} else {
				global.Logger.Debug(context.Background(), "redis set success")
			}

			//后台登录更新token，本质上就是给login接口发送请求
			req := BackendLoginReq{}
			data, err := json.MarshalIndent(&req, "", "\t")
			if err != nil {
				global.Logger.Error(context.Background(), "json parse error")
				c.Abort()
				return
			}
			request, err := http.NewRequest("POST", "http://localhost:8090/douyin/user/login?token="+accessToken, bytes.NewBuffer(data))
			if err != nil {
				global.Logger.Error(context.Background(), "login move forward error")
				c.Abort()
				return
			}
			request.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			post, err := client.Do(request)
			if post.StatusCode == 200 {
				//发送登录请求成功
				c.Set("userId", userId)
				c.Next()
				return
			} else {
				global.Logger.Error(context.Background(), "login move forward error")
				c.Abort()
				return
			}
		}
		//未过期
		userId, err := util.GetUserIDFormToken(token)
		if err != nil {
			panic(err)
		}
		c.Set("userId", userId)
		c.Next()
	}

}