package model

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	CreateBy string `json:"create_by"`
	UpdateBy string `json:"update_by"`
}
