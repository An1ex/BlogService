package model

import (
	"BlogService/pkg/app"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Tag struct {
	Model
	Name     string    `json:"name" gorm:"type:varchar(100); default:''; comment:'标签名称'"`
	State    uint8     `json:"state" gorm:"type:tinyint(3) unsigned; default:1; comment:'标签状态,0为禁用,1为启用'"`
	Articles []Article `json:"articles" gorm:"many2many:blog_article_tags; comment:'文章标签关联表'"`
}

type SwaggerTag struct {
	List  []*Tag
	Pager app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Get(db *gorm.DB) (*Tag, error) {
	var tag *Tag
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	if err := db.Model(&t).Where("id = ? AND state = ?", t.Model.ID, t.State).Preload("Articles").First(&tag).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "database: tag does not exist")
	}
	return tag, nil
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	err := db.Model(&t).Where("state = ?", t.State).Count(&count).Error
	if err != nil {
		return 0, errors.Wrap(err, "database: failed to count tags")
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
		return nil, errors.Wrap(err, "database: failed to get tag list")
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	var tag Tag
	if err := db.Where("name = ?", t.Name).First(&tag).Error; err == nil {
		return errors.New("database: tag already exists")
	}
	if err := db.Create(&t).Error; err != nil {
		return errors.Wrap(err, "database: failed to create tag")
	}
	return nil
}

func (t Tag) Update(db *gorm.DB, values map[string]interface{}) error {
	db = db.Model(&t).Where("id = ?", t.Model.ID).Take(&t)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return errors.Wrap(db.Error, "database: tag does not exist")
	}
	if err := db.Updates(values).Error; err != nil {
		return errors.Wrap(err, "database: failed to update tag")
	}
	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	db = db.Model(&t).Where("id = ?", t.Model.ID).Take(&t)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return errors.Wrap(db.Error, "database: tag does not exist")
	}
	if err := db.Select(clause.Associations).Delete(&t).Error; err != nil {
		return errors.Wrap(err, "database: failed to delete tag")
	}
	return nil
}
