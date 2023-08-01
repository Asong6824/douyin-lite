package model

import (
	"github.com/jinzhu/gorm"
	"errors"
)
type User struct {
	UserID         uint32  `json:"user_id"`
	UserName       string  `json:"user_name"`
	Password       string  `json:"password"`
	FollowingCount uint32  `json:"following_count"`
	FollowersCount  uint32  `json:"follower_count"`
}

func (u User) Register(db *gorm.DB) (uint32, error) {
	var user User
	//db.AutoMigrate(&User{})
	if err := db.Table("users").Create(u).Error; err != nil {
		return 0, err
	}
	result := db.Table("users").Where("user_name = ? AND password = ?", u.UserName, u.Password).First(&user)

	// 检查查询时是否发生错误
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 如果没有找到匹配的记录，返回错误信息
			return 0, result.Error
		}
		// 如果发生其他错误，返回错误信息
		return 0, result.Error
	}

	return user.UserID, nil
}

func (u User) Login(db *gorm.DB) (uint32, error) {
	var user User

	// 使用 GORM 的 Where 方法查询数据库中的记录
	result := db.Table("users").Where("user_name = ? AND password = ?", u.UserName, u.Password).First(&user)

	// 检查查询时是否发生错误
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 如果没有找到匹配的记录，返回错误信息
			return 0, result.Error
		}
		// 如果发生其他错误，返回错误信息
		return 0, result.Error
	}

	// 如果找到了匹配的记录，返回user的userID字段和nil错误
	return user.UserID, nil
}


func (u User) GetUser(db *gorm.DB) (*User, error) {
	if u.UserID == 0 {
		return nil, errors.New("UserID is empty")
	}

	var user User

	// 使用 GORM 的 First 方法查询数据库中的记录
	result := db.First(&user, u.UserID)

	// 检查查询时是否发生错误
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 如果没有找到匹配的记录，返回错误信息
			return nil, result.Error
		}
		// 如果发生其他错误，返回错误信息
		return nil, result.Error
	}

	// 如果找到了匹配的记录，返回查询到的用户和nil错误
	return &user, nil
}
