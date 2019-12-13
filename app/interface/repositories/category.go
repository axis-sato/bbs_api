package repositories

import (
	"github.com/c8112002/bbs_api/app/domain/eintities"
	"github.com/c8112002/bbs_api/app/interface/repositories/models"
)

type Category struct {
	DB Database
}

func (c *Category) FetchAllCategories() eintities.Categories {
	var cs models.Categories
	c.DB.Find(&cs)

	var categories eintities.Categories

	for i := range cs {
		cate := eintities.Category{ID: cs[i].ID, Name: cs[i].Name}
		categories = append(categories, cate)
	}

	return categories
}

