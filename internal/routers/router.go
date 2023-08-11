package routers

import (
	v1 "douyin/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
	"douyin/internal/middleware"
	"net/http"
	"douyin/global"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.Use(middleware.Auth())
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	user := v1.NewUser()
	video := v1.NewVideo()
	vl := v1.NewVideoLike()
	vc := v1.NewVideoComment()
	fr := v1.NewFollowRelation()
	apiv1 := r.Group("/douyin")
	{
		apiv1.GET("/feed/", video.Feed)                     //videos list
		apiv1.POST("/user/register/", user.Register)        //users create
		apiv1.POST("/user/login/", user.Login)              //users get
		apiv1.GET("/user/", user.GetUser)                   //users get
		apiv1.POST("/publish/action/", video.PublishAction) //videos create
		apiv1.GET("/publish/list/", video.PublishList)      //videos get

		apiv1.POST("/favorite/action/", vl.FavoriteAction) //video_likes create
		apiv1.GET("/favorite/list/", vl.FavoriteList)
		apiv1.POST("/comment/action/", vc.CommentAction)
		apiv1.GET("/comment/list/", vc.CommentList)

		apiv1.POST("/relation/action/", fr.FollowAction)
		apiv1.GET("/relation/follow/list/", fr.FollowList)
	}
	return r
}
