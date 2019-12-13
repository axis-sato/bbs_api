package controllers

import (
	"github.com/c8112002/bbs_api/app/infrastructure/database"
	"github.com/c8112002/bbs_api/app/interface/controllers/response"
	"github.com/c8112002/bbs_api/app/interface/repositories"
	"github.com/c8112002/bbs_api/app/usecases/category"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Category struct {
	CategoryUseCase category.UseCase
}

func NewCategoryController(db *database.Database) *Category {
	return &Category{
		CategoryUseCase: &category.Interactor{
			CategoryRepository: &repositories.Category{
				DB: db,
			},
		},
	}
}

func (category *Category) ShowAllCategories(c echo.Context) error {
	categories := category.CategoryUseCase.GetAllCategories()
	var cs response.Categories
	for i := range categories {
		c := categories[i]
		cs = append(cs, response.Category{
			ID:   c.ID,
			Name: c.Name,
		})
	}
	return c.JSON(http.StatusOK, cs)
}