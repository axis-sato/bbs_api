package handler

import (
	"github.com/c8112002/bbs_api/app/category"
	"github.com/c8112002/bbs_api/app/question"
)

type Handler struct {
	categoryStore category.Store
	questionStore question.Store
}

func NewHandler(cs category.Store, qs question.Store) *Handler {
	return &Handler{categoryStore: cs, questionStore: qs}
}
