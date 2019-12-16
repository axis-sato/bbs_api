package handler

import (
	"time"

	model2 "github.com/c8112002/bbs_api/internal/api/model"
)

type categoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type categoryListResponse struct {
	Categories []*categoryResponse `json:"categories"`
}

func NewCategoryListResponse(categories model2.Categories) *categoryListResponse {
	r := new(categoryListResponse)

	for _, c := range categories {
		cr := newCategoryResponse(c)
		r.Categories = append(r.Categories, cr)
	}

	return r
}

func newCategoryResponse(c model2.Category) *categoryResponse {
	return &categoryResponse{
		ID:   c.ID,
		Name: c.Name,
	}
}

type questionResponse struct {
	ID        int               `json:"id"`
	Title     string            `json:"title"`
	Body      string            `json:"body"`
	CreatedAt time.Time         `json:"createdAt"`
	Category  *categoryResponse `json:"category"`
}

type questionListResponse struct {
	Questions  []*questionResponse `json:"questions"`
	TotalCount int                 `json:"totalCount"`
}

func NewQuestionResponse(q *model2.Question) *questionResponse {
	return &questionResponse{
		ID:        q.ID,
		Title:     q.Title,
		Body:      q.Body,
		CreatedAt: q.CreatedAt,
		Category:  newCategoryResponse(q.Category),
	}
}

func NewQuestionListResponse(questions model2.Questions, totalCount int) *questionListResponse {
	r := new(questionListResponse)

	for _, q := range questions {
		qr := NewQuestionResponse(&q)
		r.Questions = append(r.Questions, qr)
	}

	r.TotalCount = totalCount

	return r
}