package model

import (
	"BlogService/pkg/app"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Article struct {
	Model
	Title         string `json:"title" gorm:"type:varchar(100); default:''; comment:'文章标题'"`
	Desc          string `json:"desc" gorm:"type:varchar(255); default:''; comment:'文章简述'"`
	Content       string `json:"content" gorm:"type:longtext; comment:'文章内容'"`
	CoverImageURL string `json:"cover_image_url" gorm:"type:varchar(255); default:''; comment:'封面图片地址'"`
	State         uint8  `json:"state" gorm:"type:tinyint(3) unsigned; default:1; comment:'文章状态,0为禁用,1为启用'"`
	Tags          []Tag  `json:"tags" gorm:"many2many:blog_article_tags; comment:'文章标签关联表'"`
}

type SwaggerArticle struct {
	List  []*Article
	Pager app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Get(db *gorm.DB) (*Article, error) {
	var article *Article
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	if err := db.Model(&a).Where("id = ? AND state = ?", a.Model.ID, a.State).Preload("Tags").First(&article).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "database: article does not exist")
	}
	return article, nil
}

func (a Article) Count(db *gorm.DB) (int64, error) {
	var count int64
	if a.Title != "" {
		db = db.Where("name = ?", a.Title)
	}
	err := db.Model(&a).Where("state = ?", a.State).Count(&count).Error
	if err != nil {
		return 0, errors.Wrap(err, "database: failed to count articles")
	}
	return count, nil
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if a.Title != "" {
		db = db.Where("name = ?", a.Title)
	}
	if err := db.Where("state = ?", a.State).Preload("Tags").Find(&articles).Error; err != nil {
		return nil, errors.Wrap(err, "database: failed to get article list")
	}
	return articles, nil
}

func (a Article) Create(db *gorm.DB) error {
	var article Article
	if err := db.Where("name = ?", a.Title).First(&article).Error; err == nil {
		return errors.New("database: article already exists")
	}
	if err := db.Create(&a).Error; err != nil {
		return errors.Wrap(err, "database: failed to create article")
	}
	return nil
}

func (a Article) Update(db *gorm.DB, values map[string]interface{}) error {
	db = db.Model(&a).Where("id = ?", a.Model.ID).Take(&a)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return errors.Wrap(db.Error, "database: article does not exist")
	}
	if err := db.Updates(values).Error; err != nil {
		return errors.Wrap(err, "database: failed to update article")
	}
	return nil
}

func (a Article) Delete(db *gorm.DB) error {
	db = db.Model(&a).Where("id = ?", a.Model.ID).Take(&a)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return errors.Wrap(db.Error, "database: article does not exist")
	}
	if err := db.Select(clause.Associations).Delete(&a).Error; err != nil {
		return errors.Wrap(err, "database: failed to delete article")
	}
	return nil
}
