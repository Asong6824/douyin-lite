package service

import (
	"douyin/pkg/util"
	"errors"
)

type FavoriteActionReq struct {
	Token      string `form:"token" binding:"required"`
	VideoID   uint32 `form:"video_id" binding:"required"`
	ActionType uint32 `form:"action_type" binding:"required,oneof=1 2"`
}

type FavoriteListReq struct {
	UserID uint32 `form:"user_id" binding:"required"`
	Token string  `form:"token" binding:"required"`
}

type FavoriteListResp struct {
	FavoriteList []VideoInfo
}

func (svc *Service) FavoriteAction(param FavoriteActionReq) error {
	UserID, err := util.GetUserIDFormToken(param.Token)
	if err != nil {
		return err
	}
	if param.ActionType == 1 {
		ok, err := svc.dao.IsFavorite(UserID, param.VideoID)
		if err != nil {
			return err
		}
		if ok == true {
			return errors.New("Already Favoriting")
		}
		err = svc.dao.Favorite(UserID, param.VideoID)
		if err != nil {
			return err
		}
		/*
		err = svc.dao.ModifyFollowersCount(param.ToUserID, 1) //用户的获赞数+1
		if err != nil {
			return err
		}
		*/
		err = svc.dao.ModifyVideoFavoriteCount(UserID, 1) //视频的获赞数+1
		if err != nil {
			return err
		}
		
		return nil
	} else {
		ok, err := svc.dao.IsFavorite(UserID, param.VideoID)
		if err != nil {
			return err
		}
		if ok == false {
			return errors.New("Not Favoriting")
		}
		err = svc.dao.Unfavorite(UserID, param.VideoID)
		if err != nil {
			return err
		}
		/*
		err = svc.dao.ModifyFollowersCount(param.ToUserID, 0) //用户的获赞数-1
		if err != nil {
			return err
		}
		*/
		err = svc.dao.ModifyVideoFavoriteCount(UserID, 0) //视频的获赞数-1
		if err != nil {
			return err
		}
		
		return nil
	} 
}

func (svc *Service) FavoriteList(param FavoriteListReq) (interface{}, error) {
	myuserid, err := util.GetUserIDFormToken(param.Token)
	if err != nil {
		return nil, err
	}
	//获取该用户点赞的视频id
	videoIdList, err := svc.dao.FavoriteList(param.UserID)
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

