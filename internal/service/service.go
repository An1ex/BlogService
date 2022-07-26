package service

import (
	"context"

	"BlogService/global"
	"BlogService/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(c context.Context) Service {
	svc := Service{ctx: c}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
