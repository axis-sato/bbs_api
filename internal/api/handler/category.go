package handler

import (
	"net/http"

	"github.com/c8112002/bbs_api/internal/api/utils"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/common/log"
)

func (h *Handler) Categories(c echo.Context) error {
	m, err := h.categoryStore.GetAllCategories()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(utils.NewError(err)))
	}

	return c.JSON(http.StatusOK, NewCategoryListResponse(m))
}
