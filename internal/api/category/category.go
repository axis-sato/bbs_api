package category

import (
	"github.com/c8112002/bbs_api/internal/api/model"
)

type Store interface {
	GetAllCategories() (model.Categories, error)
}
