package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuyuhome/go-gin-blog-api/global"
	"github.com/qiuyuhome/go-gin-blog-api/internal/model"
	"github.com/qiuyuhome/go-gin-blog-api/internal/routers"
	"github.com/qiuyuhome/go-gin-blog-api/pkg/logger"
	"github.com/qiuyuhome/go-gin-blog-api/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDbEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
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
	// 使用 viper 读取配置文件, 返回 viper 实例.
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	// 读取 app 配置, 配置文件中, key 是 `Server` 的 values, 赋值给 ServerSettingS 结构体中.
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	// 读取 app 配置.
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	// 读取数据库配置.
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

// 数据库配置.
func setupDbEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

// 配置日志.
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
