package service

//Article接口校验

type CountArticleRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=3,max=100"`
	Desc          string `form:"decs" binding:"min=2"`
	Content       string `form:"content" binding:"required,min=2"`
	CoverImageURL string `form:"cover_image_url" binding:"min=2,max=100"`
	CreateBy      string `form:"create_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint   `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"min=2,max=100"`
	Desc          string `form:"decs" binding:"min=2"`
	Content       string `form:"content" binding:"min=2"`
	CoverImageURL string `form:"cover_image_url" binding:"min=2,max=100"`
	UpdateBy      string `form:"update_by" binding:"required,min=2,max=100"`
	State         string `form:"state" binding:"oneof =0 1"`
}

type DeleteArticleRequest struct {
	ID uint `form:"id" binding:"required,gte=1"`
}
