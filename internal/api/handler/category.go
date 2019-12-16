package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/common/log"
)

func (h *Handler) Categories(c echo.Context) error {
	m, err := h.categoryStore.GetAllCategories()
	if err != nil {
		// TODO: エラーレスポンスを返す
		log.Error(err)
	}

	return c.JSON(http.StatusOK, NewCategoryListResponse(m))
}
