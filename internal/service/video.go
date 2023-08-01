package service

type PublishActionReq struct {
	UserID uint32 `form:"user" binding:"required,min=1,max=32"`
	Title string `form:"password" binding:"required,min=6,max=32"`
	FilePath string 
}

func (svc *Service) PublishAction(param PublishActionReq) error {
	return svc.dao.PublishAction(param.UserID, param.Title, param.FilePath)
}