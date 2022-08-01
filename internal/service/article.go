package service

import (
	"reflect"

	"BlogService/internal/model"
	"BlogService/pkg/app"
)

//Article接口校验

type GetArticleRequest struct {
	ID    uint   `form:"id" binding:"required,gte=1"`
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

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
	Desc          string `form:"desc" binding:"max=255"`
	Content       string `form:"content" binding:"required,min=2"`
	CoverImageURL string `form:"cover_image_url" binding:"min=2,max=255"`
	CreateBy      string `form:"create_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint   `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	Content       string `form:"content"`
	CoverImageURL string `form:"cover_image_url" binding:"max=255"`
	UpdateBy      string `form:"update_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint `form:"id" binding:"required,gte=1"`
}

func (svc *Service) GetArticle(param *GetArticleRequest) (*model.Article, error) {
	return svc.dao.GetArticle(param.ID, param.Title, param.State)
}

func (svc *Service) CountArticle(param *CountArticleRequest) (int64, error) {
	return svc.dao.CountArticle(param.Title, param.State)
}

func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	return svc.dao.GetArticleList(param.Title, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title, param.Desc, param.Content, param.CoverImageURL, param.State, param.CreateBy)
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	typ := reflect.TypeOf(param).Elem()
	return svc.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.Content, param.CoverImageURL, param.State, param.UpdateBy, typ)
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.ID)
}
