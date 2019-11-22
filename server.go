package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

var db *gorm.DB
var err error

func main() {

	db, err = gorm.Open("mysql", "bbs:bbspassword@tcp(localhost:3306)/bbs?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)

	defer db.Close()

	e := echo.New()

	db.SetLogger(e.Logger)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/categories", getCategories)
	e.POST("/questions", createQuestion)

	e.Logger.Fatal(e.Start(":1234"))
}

func getCategories(c echo.Context) error {
	var categories categories
	db.Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

func createQuestion(c echo.Context) error {
	q := new(question)
	if err := c.Bind(q); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	db.Create(q)
	return c.JSON(http.StatusCreated, q)

}


type category struct {
	ID   int    `json:"id" gorm:"column:id;primary_key"`
	Name string `json:"name" gorm:"column:name"`
}

type categories = []category

type question struct {
	ID int   `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Title string   `json:"title" gorm:"column:title"`
	Body string   `json:"body" gorm:"column:body"`
	CategoryId int   `json:"categoryId" gorm:"column:category_id"`
}
