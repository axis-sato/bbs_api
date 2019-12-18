package handler

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/c8112002/bbs_api/internal/api/model"
	"github.com/stretchr/testify/mock"

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

type mockCategoryStore struct {
	mock.Mock
}

func (m *mockCategoryStore) GetAllCategories() (model.Categories, error) {
	args := m.Called()

	var cl model.Categories
	if args.Get(0) != nil {
		cl = args.Get(0).(model.Categories)
	}

	return cl, args.Error(1)
}

func Testカテゴリ一覧取得_カテゴリ取得失敗(t *testing.T) {
	setup()
	defer tearDown()

	mcs := new(mockCategoryStore)
	mcs.On("GetAllCategories").Return(nil, errors.New("db error"))

	h := NewHandler(mcs, qs)

	rec, c := newRecAndContext(echo.GET, "/api/categories", nil)

	assert.NoError(t, h.Categories(c))
	assertResponse(t, rec.Result(), 500, "./testdata/category/category_list_fetch_category_error.golden")
	var er errorListResponse
	err := json.Unmarshal(rec.Body.Bytes(), &er)
	assert.NoError(t, err)
}
