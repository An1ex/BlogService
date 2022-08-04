package model

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (a Article) AppendTag(db *gorm.DB, tagNames []string, createBy string, state uint8) error {
	var tags []Tag
	for _, tagName := range tagNames {
		tag := Tag{}
		err := db.Model(&Tag{}).Where("name = ?", tagName).First(&tag).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tags = append(tags, Tag{Name: tagName, State: state, Model: Model{CreateBy: createBy}})
		} else {
			tags = append(tags, tag)
		}
	}
	if err := db.Model(&a).Association("Tags").Append(tags); err != nil {
		return errors.Wrap(err, "database: failed to append article's tag")
	}
	return nil
}

func (a Article) RemoveTag(db *gorm.DB, tagNames []string) error {
	var tags []Tag
	for _, tagName := range tagNames {
		tag := Tag{}
		if err := db.Model(&Tag{}).Where("name = ?", tagName).First(&tag).Error; err == nil {
			tags = append(tags, tag)
		}
	}
	if err := db.Model(&a).Association("Tags").Delete(tags); err != nil {
		return errors.Wrap(err, "database: failed to delete article's tag")
	}
	return nil
}
