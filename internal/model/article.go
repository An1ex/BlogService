package model

import "BlogService/pkg/app"

type Article struct {
	Model
	Title         string `json:"title" gorm:"type:varchar(100); default:''; comment:'文章标题'"`
	Desc          string `json:"desc" gorm:"type:varchar(255); default:''; comment:'文章简述'"`
	Content       string `json:"content" gorm:"type:longtext; comment:'文章内容'"`
	CoverImageURL string `json:"cover_image_url" gorm:"type:varchar(255); default:''; comment:'封面图片地址'"`
	State         uint8  `json:"state" gorm:"type:tinyint(3) unsigned; default:1; comment:'文章状态,0为禁用,1为启用'"`
}

type SwaggerArticle struct {
	List  []*Article
	Pager app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}
