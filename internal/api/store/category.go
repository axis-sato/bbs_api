package store

import (
	"github.com/c8112002/bbs_api/internal/api/model"
	"github.com/jinzhu/gorm"
)

type CategoryStore struct {
	db *gorm.DB
}

func NewCategoryStore(db *gorm.DB) *CategoryStore {
	return &CategoryStore{
		db: db,
	}
}

func (cs *CategoryStore) GetAllCategories() (model.Categories, error) {
	var m model.Categories
	err := cs.db.Find(&m).Error
	if err != nil {
		return nil, err
	}
	return m, err
}
