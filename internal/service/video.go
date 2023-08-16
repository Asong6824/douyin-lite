package service

import (
	"douyin/pkg/util"
	"time"
)

type PublishActionReq struct {
	Token string
	Title string 
	FilePath string 
}

type PublishListReq struct {
	UserID uint32 `form:"user_id" binding:"required"`
	Token string  `form:"token" binding:"required"`
}

type PublishListResp struct {
	PublishList []VideoInfo
}

type FeedReq struct {
	Token      string     `form:"token" binding:"required"`
	LatestTime string  `form:"latest_time"`
}

type FeedResp struct {
	VideoList []VideoInfo
	NextTime  time.Time
}

type VideoInfo struct {
	VideoID       uint32   `json:"id"`
	FilePath      string   `json:"play_url"`
	Title         string   `json:"title"`
	IsFavorite    bool     `json:"is_favorite"`
	UserInfo      UserInfo `json:"author"`
	FavoriteCount uint32   `json:"favorite_count"`
	CommentCount  uint32   `json:"comment_count"`
}

func (svc *Service) PublishAction(param PublishActionReq) error {
	UserID, err := util.GetUserIDFormToken(param.Token)
	if err != nil {
		return err
	}
	return svc.dao.PublishAction(UserID, param.Title, param.FilePath)
}

func (svc *Service) PublishList(param PublishListReq) (interface{}, error) {
	myuserid, err := util.GetUserIDFormToken(param.Token)
	if err != nil {
		return nil, err
	}
	//获取该用户发布的视频id
	videoIdList, err := svc.dao.GetVideoList(param.UserID)
	if err != nil {
		//global.Logger.Error(context.Background(), "get video list error")
		return nil, err
	}
	if len(videoIdList) == 0 {
		return nil, nil
	}
	//根据videoId逐一获取视频信息 获取视频信息可以再封装
    count := 0
	videoList := make([]VideoInfo, 0)
    for _, videoId := range videoIdList {
		// 达到最大数量时跳出循环
		if count >= 8 {
            break 
        }
		//获取视频信息， videos表
		rawVideoInfo, err := svc.dao.GetVideo(videoId)
		if err != nil {
            return nil, err
        }
		videoInfo := VideoInfo{
			VideoID: videoId,
			FilePath: rawVideoInfo.FilePath,
			Title: rawVideoInfo.Title,
			FavoriteCount: rawVideoInfo.FavoriteCount,
			CommentCount: rawVideoInfo.CommentCount, 
		}
		//判断是否点赞，video_likes表
		isFavorite, err := svc.dao.IsFavorite(myuserid, videoId)
		if err != nil {
            return nil, err
        }
		videoInfo.IsFavorite = isFavorite
		//获取用户信息，调用service的获取用户信息功能
		getUserParam := GetUserReq{
			UserID: param.UserID,
			Token: param.Token,
		}
        userinfo, err := svc.GetUser(getUserParam)
        if err != nil {
            return nil, err
        }
		videoInfo.UserInfo = userinfo.(UserInfo)
        videoList = append(videoList, videoInfo)
        count++
	}
	return PublishListResp{
		PublishList: videoList,
	}, nil
}

func (svc *Service) Feed(param FeedReq) (interface{}, error) {
	myuserid, err := util.GetUserIDFormToken(param.Token)
	if err != nil {
		return nil, err
	}
	var videoIdList []uint32
	if param.LatestTime != "" {
		// 使用参数中的 LatestTime
		uploadTime, err := time.Parse(time.RFC3339, param.LatestTime)
		if err != nil {
			return nil, err
		}
		videoIdList, err = svc.dao.GetVideoListByTime(uploadTime)
		if err != nil {
			return nil, err
		}
	} else {
		videoIdList, err = svc.dao.GetVideoListByTime(time.Now())
		if err != nil {
			return nil, err
		}
	}
	if len(videoIdList) == 0 {
		return nil, nil
	}
	//根据videoId逐一获取视频信息 获取视频信息可以再封装
    count := 0
	videoList := make([]VideoInfo, 0)
	var nextTime time.Time
    for _, videoId := range videoIdList {
		// 达到最大数量时跳出循环
		if count >= 8 {
            break 
        }
		//获取视频信息， videos表
		rawVideoInfo, err := svc.dao.GetVideo(videoId)
		if err != nil {
            return nil, err
        }
		videoInfo := VideoInfo{
			VideoID: videoId,
			FilePath: rawVideoInfo.FilePath,
			Title: rawVideoInfo.Title,
			FavoriteCount: rawVideoInfo.FavoriteCount,
			CommentCount: rawVideoInfo.CommentCount,
		}
		nextTime = rawVideoInfo.UploadTime
		//判断是否点赞，video_likes表
		isFavorite, err := svc.dao.IsFavorite(myuserid, videoId)
		if err != nil {
            return nil, err
        }
		videoInfo.IsFavorite = isFavorite
		//获取用户信息，调用service的获取用户信息功能
		getUserParam := GetUserReq{
			UserID: rawVideoInfo.UserID,
			Token: param.Token,
		}
        userinfo, err := svc.GetUser(getUserParam)
        if err != nil {
            return nil, err
        }
		videoInfo.UserInfo = userinfo.(UserInfo)
        videoList = append(videoList, videoInfo)
        count++
	}
	return FeedResp{
		VideoList: videoList,
		NextTime: nextTime,
	}, nil
}


