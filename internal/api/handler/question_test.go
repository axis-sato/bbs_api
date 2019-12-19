package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/c8112002/bbs_api/internal/api/model"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

func Test質問一覧取得(t *testing.T) {
	setup()
	defer tearDown()

	testcases := []struct {
		tname  string
		query  string
		golden string
	}{
		{tname: "limitなし since_idなし", query: "", golden: "no_limit_and_since_id.golden"},
		{tname: "limitあり since_idなし", query: "limit=10", golden: "limit.golden"},
		{tname: "limitなし since_idあり", query: "since_id=10", golden: "since_id.golden"},
		{tname: "limitあり since_idあり", query: "limit=10&since_id=10", golden: "limit_and_since_id.golden"},
	}

	for _, tc := range testcases {
		t.Run(tc.tname, func(t *testing.T) {
			url := fmt.Sprintf("%v?%v", "/api/questions", tc.query)
			rec, c := newRecAndContext(echo.GET, url, nil)

			assert.NoError(t, h.Questions(c))
			assertResponse(t, rec.Result(), 200, filepath.Join("./testdata/question/question_list/", tc.golden))
			var qr questionListResponse
			err := json.Unmarshal(rec.Body.Bytes(), &qr)
			assert.NoError(t, err)
		})
	}
}

func Test質問一覧取得_質問取得失敗(t *testing.T) {
	setup()
	defer tearDown()

	mqs := new(mockQuestionStore)
	mqs.On("List", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

	h := NewHandler(cs, mqs)
	rec, c := newRecAndContext(echo.GET, "/api/questions", nil)

	assert.NoError(t, h.Questions(c))
	assertResponse(t, rec.Result(), 500, "./testdata/question/question_list/question_list_error.golden")
	var er errorListResponse
	err := json.Unmarshal(rec.Body.Bytes(), &er)
	assert.NoError(t, err)
}

func Test質問作成(t *testing.T) {
	setup()
	defer tearDown()

	requestJson := `{"question": {"title": "テストタイトル", "body": "テスト本文", "categoryId": 1}}`

	rec, c := newRecAndContext(echo.POST, "/api/questions", strings.NewReader(requestJson))

	assert.NoError(t, h.CreateQuestion(c))
	assertResponse(t, rec.Result(), 200, "./testdata/question/create_question.golden")

	var qr questionResponse
	err := json.Unmarshal(rec.Body.Bytes(), &qr)
	assert.NoError(t, err)

	var qc int
	d.Model(&model.Question{}).Count(&qc)
	assert.Equal(t, 21, qc)
}

type mockQuestionStore struct {
	mock.Mock
}

func (m *mockQuestionStore) List(sinceID int, limit int) (model.Questions, error) {
	args := m.Called(sinceID, limit)

	var ql model.Questions
	if args.Get(0) != nil {
		ql = args.Get(0).(model.Questions)
	}

	return ql, args.Error(1)

}
func (m *mockQuestionStore) TotalCount() (int, error) {
	args := m.Called()

	return args.Int(0), args.Error(1)
}
func (m *mockQuestionStore) CreateQuestion(q *model.Question) error {
	args := m.Called(q)
	return args.Error(0)
}
