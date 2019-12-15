package store

import (
	"github.com/c8112002/bbs_api/app/model"
	"github.com/jinzhu/gorm"
)

type QuestionStore struct {
	db *gorm.DB
}

func NewQuestionStore(db *gorm.DB) *QuestionStore {
	return &QuestionStore{
		db: db,
	}
}

func (qs *QuestionStore) List(sinceID int, limit int) (model.Questions, error) {
	var m model.Questions

	err := qs.db.
		Where("id < ?", sinceID).
		Order("id desc").
		Preload("Category").
		Limit(limit).
		Find(&m).Error

	if err != nil {
		return nil, err
	}
	return m, err
}

func (qs *QuestionStore) TotalCount() (int, error) {
	var tc int
	err := qs.db.Model(&model.Question{}).Count(&tc).Error

	if err != nil {
		return 0, err
	}
	return tc, err
}
