package model

import "BlogService/pkg/app"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageURL string `json:"coverImageURL"`
	State         uint8
}

type SwaggerArticle struct {
	List  []*Article
	Pager app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}
