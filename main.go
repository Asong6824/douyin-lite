package main

import (
	"douyin/global"
	"douyin/internal/model"
	"douyin/internal/routers"
	"douyin/pkg/logger"
	"douyin/pkg/setting"
	"log"
	"net/http"
	"time"
	//"context"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupMysql()
	if err != nil {
		log.Fatalf("init.setupMysql err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupRedis()
	if err != nil {
		log.Fatalf("init.setupRedis err: %v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Mysql", &global.MysqlSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
    	return err
    }

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupMysql() error {
	var err error
	global.DBEngine, err = model.NewMysqlConn(global.MysqlSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupRedis() error {
	var err error
	global.Redis, err = model.NewRedisConn(global.RedisSetting)
	if err != nil {
		return err
	}

	return nil
}
