package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/common/log"
)

func (h *Handler) Questions(c echo.Context) error {
	sinceID, err := strconv.Atoi(c.QueryParam("since_id"))
	if err != nil {
		sinceID = math.MaxInt64
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	m, err := h.questionStore.List(sinceID, limit)
	if err != nil {
		// TODO: エラーレスポンスを返す
		log.Error(err)
	}

	tc, err := h.questionStore.TotalCount()
	if err != nil {
		// TODO: エラーレスポンスを返す
		log.Error(err)
	}

	return c.JSON(http.StatusOK, NewQuestionListResponse(m, tc))
}
