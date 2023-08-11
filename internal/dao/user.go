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

//ActionType 0 minus, 1 plus
func (d *Dao) ModifyFollowingCount(userid uint32, ActionType int) error {
	user := model.User{
		UserID: userid,
	}
	if ActionType == 1 {
		return user.PlusFollowingCount(d.engine)
	} else {
		return user.MinusFollowingCount(d.engine)
	}
}

func (d *Dao) ModifyFollowersCount(userid uint32, ActionType int) error {
	user := model.User{
		UserID: userid,
	}
	if ActionType == 1 {
		return user.PlusFollowersCount(d.engine)
	} else {
		return user.MinusFollowersCount(d.engine)
	}
}
