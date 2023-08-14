package v1

import (
	"github.com/gin-gonic/gin"
    "douyin/internal/service"
    "douyin/pkg/app"
    "douyin/global"
    "douyin/pkg/errcode"
	"context"
)

type VideoLike struct {}

func NewVideoLike() VideoLike {
    return VideoLike{}
}

func (vl VideoLike) FavoriteAction(c *gin.Context) {
    param := service.FavoriteActionReq{}
	response := app.NewResponse(c)
    valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errMsg := "app.BindAndValid errs: %v"
		global.Logger.Errorf(context.Background(), errMsg, errs)
		response.ToErrorResponse(errcode.InvalidParams, nil)
		return
	}
	svc := service.New(c.Request.Context())
    err := svc.FavoriteAction(param)
    if err != nil {
		errMsg :="svc.FavoriteAction err: %v"
		global.Logger.Errorf(context.Background(), errMsg, err)
		response.ToErrorResponse(errcode.ErrorFavoriteActionFail, nil)
		return
	}
	response.ToResponse(nil)
}

func (vl VideoLike) FavoriteList(c *gin.Context) {}
