package category

import (
	"github.com/c8112002/bbs_api/app/domain/eintities"
	"github.com/c8112002/bbs_api/app/usecases/repositories"
)

type Interactor struct {
	CategoryRepository repositories.Category
}

func (i *Interactor) GetAllCategories() eintities.Categories {
	return i.CategoryRepository.FetchAllCategories()
}