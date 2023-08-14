package v1

import (
	"github.com/gin-gonic/gin"
    "douyin/internal/service"
    "douyin/pkg/app"
    "douyin/global"
    "douyin/pkg/errcode"
	"context"
)

type FollowRelation struct {}

func NewFollowRelation() FollowRelation {
    return FollowRelation{}
}

func (fr FollowRelation) FollowAction(c *gin.Context) {
    param := service.FollowActionReq{}
	response := app.NewResponse(c)
    valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errMsg := "app.BindAndValid errs: %v"
		global.Logger.Errorf(context.Background(), errMsg, errs)
		response.ToErrorResponse(errcode.InvalidParams, nil)
		return
	}
	svc := service.New(c.Request.Context())
    err := svc.FollowAction(param)
    if err != nil {
		errMsg :="svc.FollowAction err: %v"
		global.Logger.Errorf(context.Background(), errMsg, err)
		response.ToErrorResponse(errcode.ErrorFollowActionFail, nil)
		return
	}
	response.ToResponse(nil)
}
//用户关注列表
func (fr FollowRelation) FollowList(c *gin.Context) {
	param := service.FollowListReq{}
	response := app.NewResponse(c)
	responseData := map[string]interface{}{
		"user_list": "",
	}
    valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errMsg := "app.BindAndValid errs: %v"
		global.Logger.Errorf(context.Background(), errMsg, errs)
		response.ToErrorResponse(errcode.InvalidParams, nil)
		return
	}
	svc := service.New(c.Request.Context())
	rawUserList, err := svc.FollowList(param)
    if err != nil {
		errMsg :="svc.FollowList err: %v"
		global.Logger.Errorf(context.Background(), errMsg, err)
		response.ToErrorResponse(errcode.ErrorFollowListFail, nil)
		return
	}
	userList := rawUserList.(service.FollowListResp)
	responseData["user_list"] = userList.UserList
	response.ToResponse(responseData)
}