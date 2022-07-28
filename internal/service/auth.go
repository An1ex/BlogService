package service

import "BlogService/internal/model"

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (svc *Service) GetAuth(param *AuthRequest) (model.Auth, error) {
	return svc.dao.GetAuth(param.AppKey, param.AppSecret)
}
