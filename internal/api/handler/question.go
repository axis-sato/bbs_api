package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/c8112002/bbs_api/internal/api/utils"

	"github.com/c8112002/bbs_api/pkg"

	model2 "github.com/c8112002/bbs_api/internal/api/model"

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
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(utils.NewError(err)))
	}

	tc, err := h.questionStore.TotalCount()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(utils.NewError(err)))
	}

	return c.JSON(http.StatusOK, NewQuestionListResponse(m, tc))
}

func (h *Handler) CreateQuestion(c echo.Context) error {
	req := new(questionCreateRequest)
	var q model2.Question

	if err := req.bind(c, &q); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, NewErrorResponse(utils.NewError(err)))
	}

	q.CreatedAt = pkg.Now()

	if err := h.questionStore.CreateQuestion(&q); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, NewErrorResponse(utils.NewError(err)))
	}

	return c.JSON(http.StatusOK, NewQuestionResponse(&q))
}
