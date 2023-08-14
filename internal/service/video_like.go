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
		err = svc.dao.ModifyVideoFavoriteCount(UserID, 1) //视频的获赞数+1
		if err != nil {
			return err
		}
		*/
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
		err = svc.dao.ModifyFollowersCount(param.ToUserID, 0) //被关注者的粉丝数-1
		if err != nil {
			return err
		}
		err = svc.dao.ModifyVideoFavoriteCount(UserID, 0) //关注者的关注数-1
		if err != nil {
			return err
		}
		*/
		return nil
	} 
}