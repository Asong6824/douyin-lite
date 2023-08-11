package service

import (
	"douyin/pkg/util"
	"errors"
)

type FollowActionReq struct {
	Token      string `form:"token" binding:"required"`
	ToUserID   uint32 `form:"to_user_id" binding:"required"`
	ActionType uint32 `form:"action_type" binding:"required,oneof=1 2"`
}

type FollowListReq struct {
	UserID uint32 `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type UserInfo = GetUserResp

type FollowListResp struct {
	UserList []UserInfo
}


 
func (svc *Service) FollowAction(param FollowActionReq) error {
	UserID, err := util.GetUserIDFormToken(param.Token)
	if err != nil {
		return err
	}
	if param.ActionType == 1 {
		ok, err := svc.dao.IsFollow(param.ToUserID, UserID)
		if err != nil {
			return err
		}
		if ok == true {
			return errors.New("Already following")
		}
		err = svc.dao.Follow(param.ToUserID, UserID)
		if err != nil {
			return err
		}
		err = svc.dao.ModifyFollowersCount(param.ToUserID, 1) //被关注者的粉丝数+1
		if err != nil {
			return err
		}
		err = svc.dao.ModifyFollowingCount(UserID, 1) //关注者的关注数+1
		if err != nil {
			return err
		}
		return nil
	} else {
		ok, err := svc.dao.IsFollow(param.ToUserID, UserID)
		if err != nil {
			return err
		}
		if ok == false {
			return errors.New("Not following")
		}
		err = svc.dao.Unfollow(param.ToUserID, UserID)
		if err != nil {
			return err
		}
		err = svc.dao.ModifyFollowersCount(param.ToUserID, 0) //被关注者的粉丝数-1
		if err != nil {
			return err
		}
		err = svc.dao.ModifyFollowingCount(UserID, 0) //关注者的关注数-1
		if err != nil {
			return err
		}
		return nil
	} 
}

func (svc *Service) FollowList(param FollowListReq) (interface{}, error) {
	userIdList, err := svc.dao.FollowList(param.UserID)
	if err != nil {
        return nil, err
    }
	if len(userIdList) == 0 {
		return nil, nil
	}
    count := 0
	userList := make([]UserInfo, 0)
    for _, userId := range userIdList {
        if count >= 11 {
            break // 达到最大数量时跳出循环
        }
		getUserParam := GetUserReq{
			UserID: userId,
			Token: param.Token,
		}
        userinfo, err := svc.GetUser(getUserParam)
        if err != nil {
            return nil, err
        }

        userList = append(userList, userinfo.(UserInfo))
        count++
    }

	return FollowListResp{
		UserList: userList,
	}, nil
}

