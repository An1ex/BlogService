package model

import (
	"BlogService/pkg/app"
	"gorm.io/gorm"
)

type Tag struct {
	Model
	Name  string `json:"name" gorm:"type:varchar(100); default:''; comment:'标签名称'"`
	State uint8  `json:"state" gorm:"type:tinyint(3) unsigned; default:1; comment:'标签状态,0为禁用,1为启用'"`
}

type SwaggerTag struct {
	List  []*Tag
	Pager app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	err := db.Model(&t).Where("state = ?", t.State).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	if err := db.Where("state = ?", t.State).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values map[string]interface{}) error {
	db = db.Model(t).Where("id = ?", t.Model.ID)
	return db.Updates(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ?", t.Model.ID).Delete(&t).Error
}
