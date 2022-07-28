package model

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Auth struct {
	Model
	AppKey    string `json:"app_key" gorm:"type:varchar(20); default:''; comment:'公钥'"`
	AppSecret string `json:"app_secret" gorm:"type:varchar(50); default:''; comment:'私钥'"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ?", a.AppKey, a.AppSecret)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, errors.Wrap(err, "database: failed to get auth")
	}
	return auth, nil
}
