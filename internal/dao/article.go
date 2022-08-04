package dao

import (
	"reflect"

	"BlogService/internal/model"
	"BlogService/pkg/app"

	"gorm.io/gorm"
)

func (d *Dao) GetArticle(id uint, title string, state uint8) (*model.Article, error) {
	article := model.Article{
		Model: model.Model{
			Model: gorm.Model{ID: id},
		},
		Title: title,
		State: state,
	}
	return article.Get(d.engine)
}

func (d *Dao) CountArticle(title string, state uint8) (int64, error) {
	article := model.Article{Title: title, State: state}
	return article.Count(d.engine)
}

func (d *Dao) GetArticleList(title string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{Title: title, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateArticle(title, desc, content, coverImageUrl string, state uint8, createdBy string, tagNames []string) error {
	var tags []model.Tag
	for _, tagName := range tagNames {
		tags = append(tags, model.Tag{Name: tagName, State: state, Model: model.Model{CreateBy: createdBy}})
	}
	article := model.Article{
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageURL: coverImageUrl,
		State:         state,
		Model: model.Model{
			CreateBy: createdBy,
		},
		Tags: tags,
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint, title, desc, content, coverImageUrl string, state uint8, updateBy string, typ reflect.Type) error {
	article := model.Article{
		Model: model.Model{
			Model: gorm.Model{ID: id},
		},
	}

	values := map[string]interface{}{
		"title":           title,
		"desc":            desc,
		"content":         content,
		"cover_image_url": coverImageUrl,
		"state":           state,
		"update_by":       updateBy,
	}

	realValues := make(map[string]interface{})

	for i := 0; i < typ.NumField(); i++ {
		tagName := typ.Field(i).Tag.Get("form")
		if v, ok := values[tagName]; ok && v != "" {
			realValues[tagName] = v
		}
	}
	return article.Update(d.engine, realValues)
}

func (d *Dao) DeleteArticle(id uint) error {
	article := model.Article{
		Model: model.Model{
			Model: gorm.Model{ID: id},
		},
	}
	return article.Delete(d.engine)
}
