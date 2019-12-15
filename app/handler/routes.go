package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Register(v1 *echo.Group) {
	categories := v1.Group("/categories")
	categories.GET("", h.Categories)

	questions := v1.Group("/questions")
	questions.GET("", h.Questions)
}
