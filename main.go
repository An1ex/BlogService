package main

import (
	"log"
	"net/http"
	"time"

	"BlogService/config"
	"BlogService/internal/dao"
	"BlogService/internal/routers"
)

func moduleInit() {
	dao.Init()
}

func main() {
	//load config
	if err := config.Init(); err != nil {
		log.Fatalf("[config] %v", err)
	}
	moduleInit()

	r := routers.NewRouter()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}
