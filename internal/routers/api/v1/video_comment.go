package v1

import (
	"github.com/gin-gonic/gin"
    "douyin/internal/service"
    "douyin/pkg/app"
    "douyin/global"
    "douyin/pkg/errcode"
	"context"
)

type VideoComment struct {}

func NewVideoComment() VideoComment {
    return VideoComment{}
}

func (vc VideoComment) Comment(c *gin.Context) {
    response := app.NewResponse(c)
    svc := service.New(c.Request.Context())
    actionType := c.Query("action_type")
    switch actionType {
    case "1":
        param := service.CommentActionReq{}
        valid, errs := app.BindAndValid(c, &param)
	    if !valid {
		    errMsg := "app.BindAndValid errs: %v"
		    global.Logger.Errorf(context.Background(), errMsg, errs)
		    response.ToErrorResponse(errcode.InvalidParams, nil)
		    return
	    }
        responseData := map[string]interface{}{
            "comment": "",
        }
        commentInfo, err := svc.CommentAction(param)
        if err != nil {
		    errMsg :="svc.CommentAction err: %v"
		    global.Logger.Errorf(context.Background(), errMsg, err)
		    response.ToErrorResponse(errcode.ErrorCommentActionFail, nil)
		    return
	    }
        responseData["comment"] = commentInfo
	    response.ToResponse(responseData)
    case "2":
        param := service.CommentDeleteReq{}
        valid, errs := app.BindAndValid(c, &param)
	    if !valid {
		    errMsg := "app.BindAndValid errs: %v"
		    global.Logger.Errorf(context.Background(), errMsg, errs)
		    response.ToErrorResponse(errcode.InvalidParams, nil)
		    return
	    }
        err := svc.CommentDelete(param)
        if err != nil {
		    errMsg :="svc.CommentDelete err: %v"
		    global.Logger.Errorf(context.Background(), errMsg, err)
		    response.ToErrorResponse(errcode.ErrorCommentActionFail, nil)
		    return
	    }
	    response.ToResponse(nil)
    default:
        // 处理无效的action_type参数
        // 返回错误信息给客户端
        errMsg := "app.BindAndValid errs"
		global.Logger.Errorf(context.Background(), errMsg)
		response.ToErrorResponse(errcode.InvalidParams, nil)
		return
    }
}

func (vc VideoComment) CommentList(c *gin.Context) {
    param := service.CommentListReq{}
	response := app.NewResponse(c)
	responseData := map[string]interface{}{
		"comment_list": "",
	}
    valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errMsg := "app.BindAndValid errs: %v"
		global.Logger.Errorf(context.Background(), errMsg, errs)
		response.ToErrorResponse(errcode.InvalidParams, nil)
		return
	}
	svc := service.New(c.Request.Context())
    rawCommentList, err := svc.CommentList(param)
    if err != nil {
		errMsg :="svc.CommentList err: %v"
		global.Logger.Errorf(context.Background(), errMsg, err)
		response.ToErrorResponse(errcode.ErrorPublishListFail, nil)
		return
	}
	commentList := rawCommentList.(service.CommentListResp)
	responseData["comment_list"] = commentList.CommentList
	response.ToResponse(responseData)
}