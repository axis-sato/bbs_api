package category

import "github.com/c8112002/bbs_api/app/domain/eintities"

type UseCase interface {
	GetAllCategories() eintities.Categories
}