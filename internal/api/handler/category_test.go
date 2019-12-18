package handler

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

func Testカテゴリ一覧取得(t *testing.T) {
	setup()
	defer tearDown()

	rec, c := newRecAndContext(echo.GET, "/api/categories", nil)

	assert.NoError(t, h.Categories(c))
	assertResponse(t, rec.Result(), 200, "./testdata/category/category_list.golden")
	var cr categoryListResponse
	err := json.Unmarshal(rec.Body.Bytes(), &cr)
	assert.NoError(t, err)
}
