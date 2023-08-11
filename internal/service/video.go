package service

import (
	"douyin/pkg/util"
)

type PublishActionReq struct {
	Token string
	Title string 
	FilePath string 
}

func (svc *Service) PublishAction(param PublishActionReq) error {
	
	UserID, err := util.GetUserIDFormToken(param.Token)
	if err != nil {
		return err
	}
	return svc.dao.PublishAction(UserID, param.Title, param.FilePath)
}