package model

import "BlogService/pkg/app"

type Tag struct {
	*Model        //嵌套匿名结构体
	Name   string `json:"name"`
	State  string `json:"state"`
}

type SwaggerTag struct {
	List  []*Tag
	Pager app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}
