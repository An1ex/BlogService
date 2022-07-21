package model

type ArticleTag struct {
	*Model
	TagID     string `json:"tagID"`
	ArticleID string `json:"articleID"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
