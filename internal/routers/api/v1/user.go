package v1

import (
	"github.com/gin-gonic/gin"
	"douyin/pkg/app"
	"douyin/global"
	"douyin/internal/service"
	"douyin/pkg/errcode"
	"context"
)
type User struct {}

func NewUser() User {
    return User{}
}

func (u User) Login(c *gin.Context) {
	param := service.UserLoginReq{}
	response := app.NewResponse(c)
	data := map[string]interface{}{
		"user_id": "",
		"token": "",
	}
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errMsg := "app.BindAndValid errs: %v"
		global.Logger.Errorf(context.Background(), errMsg, errs)
		response.ToErrorResponse(errcode.InvalidParams, data)
		return
	}
	svc := service.New(c.Request.Context())
	UserLoginResp, err := svc.Login(service.UserLoginReq{UserName: param.UserName, Password: param.Password})
	if err != nil {
		errMsg :="svc.Login err: %v"
		global.Logger.Errorf(context.Background(), errMsg, err)
		response.ToErrorResponse(errcode.ErrorUserLoginFail, data)
		return
	}
	userLoginResp, _ := UserLoginResp.(service.UserLoginResp)
	data["user_id"] = userLoginResp.UserID
	data["token"] = userLoginResp.Token
	response.ToResponse(data)
}

func (u User) Register(c *gin.Context) {
	param := service.UserRegisterReq{}
	response := app.NewResponse(c)
	data := map[string]interface{}{
		"user_id": "",
		"token": "",
	}
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errMsg := "app.BindAndValid errs: %v"
		global.Logger.Errorf(context.Background(), errMsg, errs)
		response.ToErrorResponse(errcode.InvalidParams, data)
		return
	}
	svc := service.New(c.Request.Context())
	UserRegisterResp, err := svc.Register(service.UserRegisterReq{UserName: param.UserName, Password: param.Password})
	if err != nil {
		errMsg :="svc.Register err: %v"
		global.Logger.Errorf(context.Background(), errMsg, err)
		response.ToErrorResponse(errcode.ErrorUserRegisterFail, data)
		return
	}
	userRegisterResp, _ := UserRegisterResp.(service.UserRegisterResp)
	data["user_id"] = userRegisterResp.UserID
	data["token"] = userRegisterResp.Token

	response.ToResponse(data)


}

func (u User) GetUser(c *gin.Context) {
	param := service.GetUserReq{}
	response := app.NewResponse(c)
	data := map[string]interface{}{
		"user": "",
	}
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errMsg := "app.BindAndValid errs: %v"
		global.Logger.Errorf(context.Background(), errMsg, errs)
		response.ToErrorResponse(errcode.InvalidParams, data)
		return
	}
	svc := service.New(c.Request.Context())
	GetUserResp, err := svc.GetUser(param)
	if err != nil {
		errMsg :="svc.GetUser err: %v"
		global.Logger.Errorf(context.Background(), errMsg, err)
		response.ToErrorResponse(errcode.ErrorGetUserFail, data)
		return
	}
	getUserResp, _ := GetUserResp.(service.GetUserResp)
	userData := map[string]interface{}{
		"id":             getUserResp.UserID,
		"name":           getUserResp.UserName,
		"follow_count":   getUserResp.FollowingCount,
		"follower_count": getUserResp.FollowersCount,
		"is_follow":      getUserResp.IsFollow,
	}
	data["user"] = userData
	response.ToResponse(data)
}