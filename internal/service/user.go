package service

import (
	"douyin/global"
	"douyin/pkg/util"
	"context"
	"time"
)

type UserRegisterReq struct {
	UserName string `form:"username" binding:"required,min=1,max=32"`
	Password string `form:"password" binding:"required,min=6,max=32"`
}

type UserRegisterResp struct {
	UserID uint32  `json:"user_id"`
	Token  string `json:"token"`
}

type UserLoginReq struct {
	UserName string `form:"username" binding:"required,min=1,max=32"`
	Password string `form:"password" binding:"required,min=6,max=32"`
}

type UserLoginResp struct {
	UserID uint32  `json:"user_id"`
	Token  string `json:"token"`
}

type GetUserReq struct {
	UserID uint32 `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type GetUserResp struct {
    UserID         uint32 `json:"id"`
    UserName       string `json:"name"`
    FollowingCount uint32 `json:"follow_count"`
    FollowersCount uint32 `json:"follower_count"`
    IsFollow       bool   `json:"is_follow"`
}

func (svc *Service) Register(param UserRegisterReq) (interface{}, error) {
	userId, err := svc.dao.Register(param.UserName, param.Password)
	if err != nil {
		global.Logger.Error(context.Background(), "svc.dao.Login error")
		return nil, err
	} 
	token, err := util.CreateAccessToken(userId)
	if err != nil {
		global.Logger.Error(context.Background(), "create access token error")
		return nil, err
	}

	// ref token 30d
	refreshToken, err := util.CreateRefreshToken(userId)
	if err != nil {
		global.Logger.Error(context.Background(), "create refresh token error")
		return nil, err
	}
	if err := global.Redis.Set(context.Background(), token, refreshToken, 30*24*time.Hour).Err(); err != nil {
		global.Logger.Error(context.Background(), "redis set error")
		return nil, err
	} else {
		global.Logger.Debug(context.Background(), "redis set success")
	}

	return UserRegisterResp{
		UserID: userId,
		Token:  token,
	}, nil
}

func (svc *Service) Login(param UserLoginReq) (interface{}, error) {
	userId, err := svc.dao.Login(param.UserName, param.Password)
	if err != nil {
		global.Logger.Error(context.Background(), "svc.dao.Login error")
		return nil, err
	} 
	token, err := util.CreateAccessToken(userId)
	if err != nil {
		global.Logger.Error(context.Background(), "create access token error")
		return nil, err
	}

	// ref token 30d
	refreshToken, err := util.CreateRefreshToken(userId)
	if err != nil {
		global.Logger.Error(context.Background(), "create refresh token error")
		return nil, err
	}

	//c.Header("token", token) // 不需要了

	// key: 2h token; value 30d token; key live time: 30d

	if err := global.Redis.Set(context.Background(), token, refreshToken, 30*24*time.Hour).Err(); err != nil {
		global.Logger.Error(context.Background(), "redis set error")
		return nil, err
	} else {
		global.Logger.Debug(context.Background(), "redis set success")
	}

	return UserLoginResp{
		UserID: userId,
		Token:  token,
	}, nil
	
}

func (svc *Service) GetUser(param GetUserReq) (interface{}, error) {
	myuserid, err := util.GetUserIDFormToken(param.Token)
	if err != nil {
		return nil, err
	}
	userinfo, err := svc.dao.GetUser(param.UserID)
	if err != nil {
		global.Logger.Error(context.Background(), "get user info error")
		return nil, err
	}
	isfollow, err := svc.dao.IsFollow(param.UserID, myuserid)
	if err != nil {
		global.Logger.Error(context.Background(), "confirm isfollow error")
		return nil, err
	}
	return GetUserResp{
		UserID:         userinfo.UserID,
		UserName:       userinfo.UserName,
		FollowingCount: userinfo.FollowingCount,
		FollowersCount: userinfo.FollowersCount,
		IsFollow:       isfollow,
	}, nil
}