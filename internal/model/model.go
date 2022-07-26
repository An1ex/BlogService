package model

import (
	"fmt"

	"BlogService/global"
	"BlogService/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	CreateBy string `json:"create_by" gorm:"type:varchar(100); default:''; comment:'创建人'"`
	UpdateBy string `json:"update_by" gorm:"type:varchar(100); default:''; comment:'修改人'"`
}

func NewDBEngine(database setting.DB) (*gorm.DB, error) {
	//connect database
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		database.Username,
		database.Password,
		database.Address,
		database.Database,
		database.Charset,
		database.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(database.MaxOpenConns)
	return db, nil
}

func MigrateDB() error {
	//create tables
	err := global.DBEngine.AutoMigrate(&Tag{}, &Article{}, &ArticleTag{})
	return err
}
