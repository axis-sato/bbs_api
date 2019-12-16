package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/c8112002/bbs_api/internal/api/router"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

func Testカテゴリ一覧取得(t *testing.T) {
	setup()
	defer tearDown()

	e := router.New()
	req := httptest.NewRequest(echo.GET, "/api/categories", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, h.Categories(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var cr categoryListResponse
		err := json.Unmarshal(rec.Body.Bytes(), &cr)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(cr.Categories))
	}
}
