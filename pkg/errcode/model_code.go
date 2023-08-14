package errcode

var (
	ErrorUserRegisterFail   = NewError(20010001, "用户注册失败")
	ErrorUserLoginFail      = NewError(20010002, "用户登录失败")
	ErrorGetUserFail        = NewError(20010003, "获取用户信息失败")
	ErrorPublishActionFail  = NewError(20020001, "上传视频失败")
	ErrorPublishListFail    = NewError(20020002, "获取发布列表失败")
	ErrorFollowActionFail   = NewError(20030001, "关注操作失败")
	ErrorFollowListFail     = NewError(20030002, "获取关注列表失败")
	ErrorFavoriteActionFail = NewError(20040002, "视频点赞失败")
)