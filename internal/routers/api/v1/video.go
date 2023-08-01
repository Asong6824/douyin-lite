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

func (v Video) PublishList(c *gin.Context) {}
func (v Video) PublishAction(c *gin.Context) {
    response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("data")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := upload.TypeVideo
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorPublishActionFail.WithDetails(err.Error()))
		return
	}

    err = svc.PublishAction(service.PublishActionReq{
        UserID: 1, 
        Title: convert.StrTo(c.PostForm("title")).String(),
        FilePath: fileInfo.AccessUrl,
    })

    if err != nil {
		errMsg :="svc.PublishAction err: %v"
		global.Logger.Errorf(context.Background(), errMsg, err)
		response.ToErrorResponse(errcode.ErrorPublishActionFail)
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
func (v Video) Feed(c *gin.Context) {}