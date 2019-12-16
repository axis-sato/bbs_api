package handler

import (
	category2 "github.com/c8112002/bbs_api/internal/api/category"
	question2 "github.com/c8112002/bbs_api/internal/api/question"
)

type Handler struct {
	categoryStore category2.Store
	questionStore question2.Store
}

func NewHandler(cs category2.Store, qs question2.Store) *Handler {
	return &Handler{categoryStore: cs, questionStore: qs}
}
