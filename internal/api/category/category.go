package category

import (
	model2 "github.com/c8112002/bbs_api/internal/api/model"
)

type Store interface {
	GetAllCategories() (model2.Categories, error)
}
