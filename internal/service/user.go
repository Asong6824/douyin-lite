package service



type UserRegisterReq struct {
	UserName string `form:"username" binding:"required,min=1,max=32"`
	Password string `form:"password" binding:"required,min=6,max=32"`
}

type UserRegisterResp struct {
	UserID uint32  `json:"user_id"`
	Token  string `json:"token"`
}

type UserLoginReq struct {
	UserName string `form:"username" binding:"required,min=1,max=32"`
	Password string `form:"password" binding:"required,min=6,max=32"`
}

type UserLoginResp struct {
	UserID uint32  `json:"user_id"`
	Token  string `json:"token"`
}

func (svc *Service) Register(param UserRegisterReq) (uint32, error) {
	return svc.dao.Register(param.UserName, param.Password)
}

func (svc *Service) Login(param UserLoginReq) (uint32, error) {
	return svc.dao.Login(param.UserName, param.Password)
}