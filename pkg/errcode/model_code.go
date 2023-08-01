package errcode

var (
	ErrorUserRegisterFail = NewError(20010001, "用户注册失败")
	ErrorUserLoginFail = NewError(20010002, "用户登录失败")
	ErrorUserGetFail = NewError(20010003, "获取用户信息失败")
	ErrorPublishActionFail = NewError(20020001, "上传视频失败")
)