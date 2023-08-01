package dao

import (
	"douyin/internal/model"
)

func (d *Dao) Register(username string, password string) (uint32, error) {
	user := model.User{
		UserName: username,
		Password: password,
	}
	return user.Register(d.engine)
}

func (d *Dao) Login(username string, password string) (uint32, error) {
	user := model.User{
		UserName: username,
		Password: password,
	}
	return user.Login(d.engine)
}

func (d *Dao) GetUser(userid uint32) (*model.User, error) {
	user := model.User{
		UserID: userid,
	}
	return user.GetUser(d.engine)
}