package model

import (
	"github.com/jinzhu/gorm"
	"errors"
)
type User struct {
	UserID         uint32  `json:"user_id"`
	UserName       string  `json:"user_name"`
	Password       string  `json:"password"`
	FollowingCount uint32  `json:"following_count"` //关注数
	FollowersCount uint32  `json:"follower_count"`  //粉丝数
}

func (u User) Register(db *gorm.DB) (uint32, error) {
	if err := db.Table("users").Model(&User{}).Create(&u).Error; err != nil {
		return 0, err
	}
	// 执行查询获取最后插入的自增主键值
	var id uint32
	result := db.Raw("SELECT LAST_INSERT_ID() as id").Scan(&id)
	if result.Error != nil {
		return 0, result.Error
	}
	return id, nil
}

func (u User) Login(db *gorm.DB) (uint32, error) {
	var user User

	// 使用 GORM 的 Where 方法查询数据库中的记录
	result := db.Table("users").
		Where("user_name = ? AND password = ?", u.UserName, u.Password).
		First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, result.Error
		}
		return 0, result.Error
	}

	return user.UserID, nil
}

func (u User) GetUser(db *gorm.DB) (*User, error) {
	var user User

	result := db.Table("users").Where("user_id = ?", u.UserID).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &user, nil
}

func (u User) PlusFollowingCount(db *gorm.DB) error {
	err := db.Table("users").
		Where("user_id = ?", u.UserID).
		Model(&u).
		UpdateColumn("following_count", gorm.Expr("following_count + ?", 1)).Error 
	if err != nil { 
		return err 
	} 
	return nil
}

func (u User) MinusFollowingCount(db *gorm.DB) error {
	err := db.Table("users").
		Where("user_id = ?", u.UserID).
		Model(&u).
		UpdateColumn("following_count", gorm.Expr("following_count - ?", 1)).Error 
	if err != nil { 
		return err 
	} 
	return nil
}

func (u User) PlusFollowersCount(db *gorm.DB) error {
	err := db.Table("users").
		Where("user_id = ?", u.UserID).
		Model(&u).
		UpdateColumn("followers_count", gorm.Expr("followers_count + ?", 1)).Error 
	if err != nil { 
		return err 
	} 
	return nil
}

func (u User) MinusFollowersCount(db *gorm.DB) error {
	err := db.Table("users").
		Where("user_id = ?", u.UserID).
		Model(&u).
		UpdateColumn("followers_count", gorm.Expr("followers_count - ?", 1)).Error 
	if err != nil { 
		return err 
	} 
	return nil
}
