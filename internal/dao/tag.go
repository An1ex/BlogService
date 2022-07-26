package dao

import (
	"BlogService/internal/model"
	"BlogService/pkg/app"

	"gorm.io/gorm"
)

func (d *Dao) CountTag(name string, state uint8) (int64, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: model.Model{
			CreateBy: createdBy,
		},
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint, name string, state uint8, updateBy string) error {
	tag := model.Tag{
		Model: model.Model{
			Model: gorm.Model{ID: id},
		},
	}

	values := map[string]interface{}{
		"state":     state,
		"update_by": updateBy,
	}
	if name != "" {
		values["name"] = name
	}
	return tag.Update(d.engine, values)
}

func (d Dao) DeleteTag(id uint) error {
	tag := model.Tag{
		Model: model.Model{
			Model: gorm.Model{ID: id},
		},
	}
	return tag.Delete(d.engine)
}