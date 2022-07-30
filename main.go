package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"BlogService/global"
	"BlogService/internal/model"
	"BlogService/internal/routers"
	"BlogService/pkg/flag"
	"BlogService/pkg/logger"
	"BlogService/pkg/setting"
	"BlogService/pkg/tracer"
	"github.com/gin-gonic/gin"
)

var (
	port          string
	runMode       string
	configPath    string
	isVersion     bool
	buildTime     string
	buildVersion  string
	commitID      string
	serviceName   string //jaeger setting
	agentHostPort string //jaeger setting
)

func init() {
	err := flag.InitFlag(&port, &runMode, &configPath, &isVersion, &serviceName, &agentHostPort)
	if err != nil {
		log.Fatalf("[flag] %v", err)
	}

	err = setting.InitSetting(port, runMode, configPath)
	if err != nil {
		log.Fatalf("[configs] %v", err)
	}

	err = logger.InitLogger()
	if err != nil {
		log.Fatalf("[logger] %v", err)
	}

	//前提：create database blog_service;
	err = model.InitDBEngine()
	if err != nil {
		log.Fatalf("[database] %v", err)
	}

	err = tracer.InitTracer(serviceName, agentHostPort)
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
	if isVersion {
		fmt.Printf("build_time: %s\n", buildTime)       // "-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S`"
		fmt.Printf("build_version: %s\n", buildVersion) // "-X main.buildVersion=1.0.0"
		fmt.Printf("commit_id: %s\n", commitID)         // "-X main.commitID=`git rev-parse HEAD`"
	}
	gin.SetMode(global.Config.Server.RunMode)
	r := routers.NewRouter()

	s := &http.Server{
		Addr:         ":" + global.Config.Server.HttpPort,
		Handler:      r,
		ReadTimeout:  global.Config.Server.ReadTimeout,
		WriteTimeout: global.Config.Server.WriteTimeout,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			global.Logger.Fatalf("[server] %v", err)
		}
	}()
	quit := make(chan os.Signal)
	//将进程收到的结束程序signal转发给channel
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Println("[server] shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		global.Logger.Fatalf("[server] forced to shutting down: %v", err)
	}
	global.Logger.Println("[server] server exiting")
}
