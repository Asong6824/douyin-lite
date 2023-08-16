package service

import (
	"douyin/pkg/util"
	"errors"
	"time"
)

type CommentActionReq struct {
	Token   string  `form:"token" binding:"required"`
	VideoID uint32  `form:"video_id" binding:"required"`
	Content string  `form:"comment_text" binding:"required"`
}

type CommentActionResp struct {
	CommentID uint32  `json:"id"`
	User UserInfo     `json:"user"`
	Content string    `json:"content"`
	CreateDate string `json:"create_date"`
}

type CommentInfo = CommentActionResp

type CommentDeleteReq struct {
	Token     string  `form:"token" binding:"required"`
	VideoID   uint32  `form:"video_id" binding:"required"`
	CommentID uint32  `form:"comment_id" binding:"required"`
}

type CommentListReq struct {
	Token     string  `form:"token" binding:"required"`
	VideoID   uint32  `form:"video_id" binding:"required"`
}

type CommentListResp struct {
	CommentList []CommentInfo
}

func (svc *Service) CommentAction(param CommentActionReq) (interface{}, error) {
	userId, err := util.GetUserIDFormToken(param.Token)
	if err != nil {
		return nil, err
	}
	commentTime := time.Now()
	commentId, err := svc.dao.CommentAction(userId, param.VideoID, param.Content, commentTime)
	if err != nil {
		return nil,err
	}
	userInfo, err := svc.GetUser(GetUserReq{Token: param.Token, UserID: userId})
	if err != nil {
		return nil,err
	}
	return CommentActionResp{
		CommentID: commentId,
		User: userInfo.(UserInfo),
		Content: param.Content,
		CreateDate: commentTime.Format("01-02"),
	}, nil
}

func (svc *Service) CommentDelete(param CommentDeleteReq) error {
	userId, err := util.GetUserIDFormToken(param.Token)
	if err != nil {
		return err
	}
	ok, err := svc.dao.IsCommented(param.CommentID, userId)
	if err != nil {
		return err
	}
	if ok == false {
		return errors.New("No comments yet")
	}
	return svc.dao.CommentDelete(param.CommentID)
}

func (svc *Service) CommentList(param CommentListReq) (interface{}, error) {
	//获取视频的评论id表
	commentIdList, err := svc.dao.CommentIdList(param.VideoID)
	if err != nil {
		//global.Logger.Error(context.Background(), "get video list error")
		return nil, err
	}
	if len(commentIdList) == 0 {
		return nil, nil
	}
	//根据commentId逐一获取评论信息
    count := 0
	commentList := make([]CommentInfo, 0)
    for _, commentId := range commentIdList {
		// 达到最大数量时跳出循环
		if count >= 4 {
            break 
        }
		//获取评论信息， videos表
		rawCommentInfo, err := svc.dao.GetComment(commentId)
		if err != nil {
            return nil, err
        }
		commentInfo := CommentInfo{
			CommentID: commentId,
			Content: rawCommentInfo.Content,
			CreateDate: rawCommentInfo.CommentTime.Format("01-02"),
		}
		//获取用户信息，调用service的获取用户信息功能
		getUserParam := GetUserReq{
			UserID: rawCommentInfo.UserID,
			Token: param.Token,
		}
        userinfo, err := svc.GetUser(getUserParam)
        if err != nil {
            return nil, err
        }
		commentInfo.User = userinfo.(UserInfo)
        commentList = append(commentList, commentInfo)
        count++
	}
	return CommentListResp{
		CommentList: commentList,
	}, nil
}