package handler

import (
	"github.com/c8112002/bbs_api/internal/api/model"
	"github.com/labstack/echo/v4"
)

type questionCreateRequest struct {
	Question struct {
		Title      string `json:"title" validate:"required,min=1,max=255"`
		Body       string `json:"body" validate:"required,min=1,max=5000"`
		CategoryID int    `json:"categoryId" validate:"required,numeric"`
	} `json:"question"`
}

func (r *questionCreateRequest) bind(c echo.Context, q *model.Question) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	q.Title = r.Question.Title
	q.Body = r.Question.Body
	q.CategoryID = r.Question.CategoryID

	return nil
}
