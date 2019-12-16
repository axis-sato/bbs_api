package handler

import (
	"log"

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
	d  *gorm.DB
	h  *Handler
	cs category2.Store
	qs question2.Store
	_  *echo.Echo
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
