package dao

import (
	"fmt"

	"BlogService/config"
	"BlogService/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connect() *gorm.DB {
	//connect to database
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.SQL.Username,
		config.SQL.Password,
		config.SQL.Address,
		config.SQL.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func Init() {
	//前提：create database blog_service;
	DB = connect()
	//create tables
	err := DB.AutoMigrate(&model.Tag{}, &model.Article{}, &model.ArticleTag{})
	if err != nil {
		panic(err)
	}
}
