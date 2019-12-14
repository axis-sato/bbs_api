package handler

import (
	"github.com/c8112002/bbs_api/app/model"
)

type categoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type categoryListResponse struct {
	Categories []*categoryResponse `json:"categories"`
}

func NewCategoryListResponse(categories model.Categories) *categoryListResponse {
	r := new(categoryListResponse)

	for _, c := range categories {
		cr := &categoryResponse{
			ID:   c.ID,
			Name: c.Name,
		}
		r.Categories = append(r.Categories, cr)
	}

	return r
}
