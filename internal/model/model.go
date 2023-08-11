package model

import (
	"douyin/global"
	"douyin/pkg/setting"
	"fmt"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewMysqlConn(databaseSetting *setting.MysqlSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

func NewRedisConn(databaseSetting *setting.RedisSettingS) (*redis.Client, error) {
	Redis := redis.NewClient(&redis.Options{
		Addr:        databaseSetting.Host,
		Password:    databaseSetting.Password,
		DB:          databaseSetting.DB,
		IdleTimeout: databaseSetting.IdleTimeout,
	})
	if _, err := Redis.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return Redis, nil
}
