package model

import (
	"fmt"

	"BlogService/global"
	"BlogService/pkg/setting"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	CreateBy string `json:"create_by" gorm:"type:varchar(100); default:''; comment:'创建人'"`
	UpdateBy string `json:"update_by" gorm:"type:varchar(100); default:''; comment:'修改人'"`
}

// NewDBEngine 连接数据库软件并创建一个数据库引擎
func NewDBEngine(database setting.DB) (*gorm.DB, error) {
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
		return nil, errors.Wrap(err, "database: unable to connect to the database")
	}
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(database.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(database.MaxOpenConns)
	return db, nil
}

// MigrateDB 迁移数据表
func MigrateDB() error {
	err := global.DBEngine.AutoMigrate(&Tag{}, &Article{}, &ArticleTag{}, &Auth{})
	return errors.Wrap(err, "database: failed to migration schema")
}
