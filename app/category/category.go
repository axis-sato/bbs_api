package category

import (
	"github.com/c8112002/bbs_api/app/model"
)

type Store interface {
	GetAllCategories() (model.Categories, error)
}
