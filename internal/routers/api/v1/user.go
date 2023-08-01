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
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errMsg := "app.BindAndValid errs: %v"
		global.Logger.Errorf(context.Background(), errMsg, errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	userid, err := svc.Login(service.UserLoginReq{UserName: param.UserName, Password: param.Password})
	if err != nil {
		errMsg :="svc.Register err: %v"
		global.Logger.Errorf(context.Background(), errMsg, err)
		response.ToErrorResponse(errcode.ErrorUserLoginFail)
		return
	}
	data := service.UserLoginResp{
		UserID: userid,
		Token: "log-token",
	}
	response.ToResponse(data)
}

func (u User) Register(c *gin.Context) {
	param := service.UserRegisterReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errMsg := "app.BindAndValid errs: %v"
		global.Logger.Errorf(context.Background(), errMsg, errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	//pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	userid, err := svc.Register(service.UserRegisterReq{UserName: param.UserName, Password: param.Password})
	if err != nil {
		errMsg :="svc.Register err: %v"
		global.Logger.Errorf(context.Background(), errMsg, err)
		response.ToErrorResponse(errcode.ErrorUserRegisterFail)
		return
	}

	data := service.UserRegisterResp{
		UserID: userid,
		Token: "dafaqa",//token
	}
	response.ToResponse(data)


}

func (u User) GetUser(c *gin.Context) {}