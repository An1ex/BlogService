package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
	"time"

	"BlogService/global"
	"BlogService/internal/model"
	"BlogService/internal/routers"
	"BlogService/pkg/logger"
	"BlogService/pkg/setting"
	"BlogService/pkg/tracer"

	"github.com/gin-gonic/gin"
)

var (
	port       string
	runMode    string
	configPath string
)

func init() {
	err := initFlag()
	if err != nil {
		log.Fatalf("[flag] %v", err)
	}

	err = loadConfig()
	if err != nil {
		log.Fatalf("[configs] %v", err)
	}

	err = initLogger()
	if err != nil {
		log.Fatalf("[logger] %v", err)
	}

	//前提：create database blog_service;
	err = initDBEngine()
	if err != nil {
		log.Fatalf("[database] %v", err)
	}

	err = initTracer()
	if err != nil {
		log.Fatalf("[tracer] %v", err)
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
	s, err := setting.NewSetting(strings.Split(configPath, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.Config.Server)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.Config.App)
	if err != nil {
		return err
	}
	err = s.ReadSection("DataBase", &global.Config.BD)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.Config.JWT)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.Config.Email)
	if err != nil {
		return err
	}
	err = s.ReadSection("Limiter", &global.Config.Limiter)
	if err != nil {
		return err
	}

	if port != "" {
		global.Config.Server.HttpPort = port
	}
	if runMode != "" {
		global.Config.Server.RunMode = runMode
	}
	global.Config.Server.ReadTimeout *= time.Second
	global.Config.Server.WriteTimeout *= time.Second
	global.Config.JWT.Expire *= time.Second
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

func initTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", "localhost:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

func initFlag() error {
	flag.StringVar(&port, "port", "", " 启动端口")
	flag.StringVar(&runMode, "mode", "", " 启动模式")
	flag.StringVar(&configPath, "path", "configs/", "指定要使用的配置文件路径")
	flag.Parse()
	return nil
}
