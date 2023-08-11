package global

import (
	"douyin/pkg/logger"
	"douyin/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	MysqlSetting    *setting.MysqlSettingS
	RedisSetting    *setting.RedisSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
)
