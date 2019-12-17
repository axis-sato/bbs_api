package handler

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/c8112002/bbs_api/pkg"

	"github.com/stretchr/testify/assert"

	"github.com/c8112002/bbs_api/internal/api/model"

	"github.com/c8112002/bbs_api/internal/api/category"
	"github.com/c8112002/bbs_api/internal/api/db"
	"github.com/c8112002/bbs_api/internal/api/question"
	"github.com/c8112002/bbs_api/internal/api/router"
	"github.com/c8112002/bbs_api/internal/api/store"

	"github.com/labstack/echo/v4"

	"github.com/jinzhu/gorm"
)

var (
	d      *gorm.DB
	h      *Handler
	cs     category.Store
	qs     question.Store
	_      *echo.Echo
	update = flag.Bool("update", false, "update .golden files")
)

func setup() {
	d = db.TestDB()
	cs = store.NewCategoryStore(d)
	qs = store.NewQuestionStore(d)
	h = NewHandler(cs, qs)
	_ = router.New()

	pkg.Freeze(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))

	if err := db.Migrate(d); err != nil {
		log.Fatal(err)
	}

	if err := loadFixtures(); err != nil {
		log.Fatal(err)
	}
}

func tearDown() {
	_ = d.Close()
	if err := db.DropTestDB(); err != nil {
		log.Fatal(err)
	}
}

func loadFixtures() error {
	var cl []model.Category
	for i := 1; i <= 2; i++ {
		c := model.Category{
			ID:   i,
			Name: fmt.Sprintf("カテゴリ%d", i),
		}
		if err := d.Create(&c).Error; err != nil {
			return nil
		}
		cl = append(cl, c)
	}

	for i := 1; i <= 20; i++ {
		c := cl[i%len(cl)]
		q := model.Question{
			ID:         i,
			Title:      fmt.Sprintf("タイトル%d", i),
			Body:       fmt.Sprintf("本文%d", i),
			CreatedAt:  pkg.Now(),
			CategoryID: c.ID,
			Category:   c,
		}
		if err := d.Create(&q).Error; err != nil {
			return nil
		}
	}

	return nil
}

func newRecAndContext(method, target string, body io.Reader) (*httptest.ResponseRecorder, echo.Context) {
	e := router.New()
	req := httptest.NewRequest(method, target, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return rec, c
}

func assertResponse(t *testing.T, res *http.Response, code int, path string) {
	t.Helper()

	assertResponseHeader(t, res, code)
	assertResponseBody(t, res, path)
}

func assertResponseHeader(t *testing.T, res *http.Response, code int) {
	t.Helper()

	if code != res.StatusCode {
		t.Errorf("expected status code is '%d',\n but actual given code is '%d'", code, res.StatusCode)
	}

	if expected := "application/json; charset=UTF-8"; res.Header.Get("Content-Type") != expected {
		t.Errorf("unexpected response Content-Type,\n expected: %#v,\n but given #%v", expected, res.Header.Get("Content-Type"))
	}
}

func assertResponseBody(t *testing.T, res *http.Response, path string) {
	t.Helper()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("unexpected error by ioutil.ReadAll() '%#v'", err)
	}

	var actual bytes.Buffer
	err = json.Indent(&actual, body, "", "  ")
	if err != nil {
		t.Fatalf("unexpected error by json.Indent '%#v'", err)
	}

	if *update {
		updateGoldenFile(t, actual, path)
	}

	rs := getStringFromTestFile(t, path)

	assert.JSONEq(t, rs, actual.String())
}

func getStringFromTestFile(t *testing.T, path string) string {
	t.Helper()

	bt, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("unexpected error while opening file '%#v'", err)
	}
	return string(bt)
}

func updateGoldenFile(t *testing.T, actual bytes.Buffer, path string) {
	t.Helper()

	t.Log("update golden file")
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0744); err != nil {
		t.Fatalf("failed to make the directory for golden file: %s", err)
	}
	if err := ioutil.WriteFile(path, actual.Bytes(), 0644); err != nil {
		t.Fatalf("failed to update golden file: %s", err)
	}
}
