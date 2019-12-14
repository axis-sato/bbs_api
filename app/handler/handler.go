package handler

import (
	"github.com/c8112002/bbs_api/app/category"
)

type Handler struct {
	categoryStore category.Store
}

func NewHandler(cs category.Store) *Handler {
	return &Handler{categoryStore: cs}
}
