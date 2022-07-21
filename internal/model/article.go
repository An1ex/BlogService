package model

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageURL string `json:"coverImageURL"`
	State         uint8
}

func (a Article) TableName() string {
	return "blog_article"
}
