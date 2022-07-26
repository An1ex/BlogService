package model

type ArticleTag struct {
	*Model
	TagID     uint `json:"tag_id"`
	ArticleID uint `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
