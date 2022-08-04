package dao

import (
	"BlogService/internal/model"
	"gorm.io/gorm"
)

func (d *Dao) AppendTag(id uint, tagNames []string, createBy string, state uint8) error {
	article := model.Article{
		Model: model.Model{
			Model: gorm.Model{ID: id},
		},
	}
	return article.AppendTag(d.engine, tagNames, createBy, state)
}

func (d *Dao) RemoveTag(id uint, tagNames []string) error {
	article := model.Article{
		Model: model.Model{
			Model: gorm.Model{ID: id},
		},
	}
	return article.RemoveTag(d.engine, tagNames)
}
