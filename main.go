package main

import (
	"log"
	"net/http"
	"time"

	"BlogService/global"
	"BlogService/internal/model"
	"BlogService/internal/routers"
	"BlogService/pkg/logger"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

func init() {
	err := loadConfig()
	if err != nil {
		log.Fatalf("[config] %v", err)
	}

	err = initLogger()
	if err != nil {
		log.Fatalf("[database] %v", err)
	}

	//前提：create database blog_service;
	err = initDBEngine()
	if err != nil {
		log.Fatalf("[database] %v", err)
	}
}

// @title           博客平台
// @version         1.0
// @description     一个完整的博客后端
// @termsOfService  https://github.com/An1ex/BlogService
// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
func main() {
	gin.SetMode(global.Config.Server.RunMode)
	r := routers.NewRouter()

	s := &http.Server{
		Addr:         ":" + global.Config.Server.HttpPort,
		Handler:      r,
		ReadTimeout:  global.Config.Server.ReadTimeout,
		WriteTimeout: global.Config.Server.WriteTimeout,
	}

	s.ListenAndServe()
}

func loadConfig() error {
	_, err := toml.DecodeFile("config/config.toml", &global.Config)
	if err != nil {
		return err
	}
	global.Config.Server.ReadTimeout *= time.Second
	global.Config.Server.WriteTimeout *= time.Second
	return nil
}

func initDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.Config.BD)
	if err != nil {
		return err
	}
	err = model.MigrateDB()
	if err != nil {
		return err
	}
	return nil
}

func initLogger() error {
	var err error
	global.Logger, err = logger.NewLogger()
	if err != nil {
		return err
	}
	return nil
}
