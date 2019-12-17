package handler

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/c8112002/bbs_api/internal/api/model"

	category2 "github.com/c8112002/bbs_api/internal/api/category"
	db2 "github.com/c8112002/bbs_api/internal/api/db"
	question2 "github.com/c8112002/bbs_api/internal/api/question"
	router2 "github.com/c8112002/bbs_api/internal/api/router"
	store2 "github.com/c8112002/bbs_api/internal/api/store"

	"github.com/labstack/echo/v4"

	"github.com/jinzhu/gorm"
)

var (
	d      *gorm.DB
	h      *Handler
	cs     category2.Store
	qs     question2.Store
	_      *echo.Echo
	update = flag.Bool("update", false, "update .golden files")
)

func setup() {
	d = db2.TestDB()
	cs = store2.NewCategoryStore(d)
	qs = store2.NewQuestionStore(d)
	h = NewHandler(cs, qs)
	_ = router2.New()

	if err := db2.Migrate(d); err != nil {
		log.Fatal(err)
	}

	if err := loadFixtures(); err != nil {
		log.Fatal(err)
	}
}

func tearDown() {
	_ = d.Close()
	if err := db2.DropTestDB(); err != nil {
		log.Fatal(err)
	}
}

func loadFixtures() error {
	c1 := model.Category{
		ID:   1,
		Name: "カテゴリ1",
	}
	if err := d.Create(&c1).Error; err != nil {
		return nil
	}
	c2 := model.Category{
		ID:   2,
		Name: "カテゴリ2",
	}
	if err := d.Create(&c2).Error; err != nil {
		return nil
	}

	return nil
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
