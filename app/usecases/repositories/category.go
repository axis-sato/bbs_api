package repositories

import "github.com/c8112002/bbs_api/app/domain/eintities"

type Category interface {
	FetchAllCategories() eintities.Categories
}
