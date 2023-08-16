package v1

import (
	"github.com/gin-gonic/gin"
	"douyin/pkg/app"
	"douyin/global"
	"douyin/internal/service"
	"douyin/pkg/errcode"
    "douyin/pkg/upload"
    "douyin/pkg/convert"
	"context"
)

type Video struct {}

func NewVideo() Video {
    return Video{}
}

func (v Video) PublishList(c *gin.Context) {
	param := service.PublishListReq{}
	response := app.NewResponse(c)
	responseData := map[string]interface{}{
		"video_list": "",
	}
    valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errMsg := "app.BindAndValid errs: %v"
		global.Logger.Errorf(context.Background(), errMsg, errs)
		response.ToErrorResponse(errcode.InvalidParams, nil)
		return
	}
	svc := service.New(c.Request.Context())
	rawVideoList, err := svc.PublishList(param)
    if err != nil {
		errMsg :="svc.PublishList err: %v"
		global.Logger.Errorf(context.Background(), errMsg, err)
		response.ToErrorResponse(errcode.ErrorPublishListFail, nil)
		return
	}
	videoList := rawVideoList.(service.PublishListResp)
	responseData["video_list"] = videoList.PublishList
	response.ToResponse(responseData)
}

func (v Video) PublishAction(c *gin.Context) {
    response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("data")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams, nil)
		return
	}

	fileType := upload.TypeVideo
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams, nil)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorPublishActionFail, nil)
		return
	}
    err = svc.PublishAction(service.PublishActionReq{
        Token: convert.StrTo(c.PostForm("token")).String(), 
        Title: convert.StrTo(c.PostForm("title")).String(),
        FilePath: fileInfo.AccessUrl,
    })

    if err != nil {
		errMsg :="svc.PublishAction err: %v"
		global.Logger.Errorf(context.Background(), errMsg, err)
		response.ToErrorResponse(errcode.ErrorPublishActionFail, nil)
		return
	}

	response.ToResponse(nil)
}

func (v Video) Feed(c *gin.Context) {
	param := service.FeedReq{}
	response := app.NewResponse(c)
	responseData := map[string]interface{}{
		"video_list": "",
		"next_time" : "",
	}
    valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errMsg := "app.BindAndValid errs: %v"
		global.Logger.Errorf(context.Background(), errMsg, errs)
		response.ToErrorResponse(errcode.InvalidParams, nil)
		return
	}
	svc := service.New(c.Request.Context())
	rawVideoList, err := svc.Feed(param)
    if err != nil {
		errMsg :="svc.PublishList err: %v"
		global.Logger.Errorf(context.Background(), errMsg, err)
		response.ToErrorResponse(errcode.ErrorFeedFail, nil)
		return
	}
	feedResp := rawVideoList.(service.FeedResp)
	responseData["video_list"] = feedResp.VideoList
	responseData["next_time"]  = feedResp.NextTime.Format("01-02")
	response.ToResponse(responseData)
}