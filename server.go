package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"math"
	"net/http"
	"strconv"
)

var db *gorm.DB
var err error

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func validateCategoryId(fl validator.FieldLevel) bool {
	categoryId := fl.Field().Int()
	c := new(category)
	db.First(&c, categoryId)
	return c.ID == int(categoryId)
}

func main() {

	db, err = gorm.Open("mysql", "bbs:bbspassword@tcp(localhost:3306)/bbs?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)

	defer func() {
		_ = db.Close()
	}()

	e := echo.New()

	v:= validator.New()
	_ = v.RegisterValidation("categoryId", validateCategoryId)
	e.Validator = &CustomValidator{validator: v}

	db.SetLogger(e.Logger)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/categories", getCategories)
	e.GET("/questions", getQuestions)
	e.POST("/questions", createQuestion)

	e.Logger.Fatal(e.Start(":1234"))
}

func getCategories(c echo.Context) error {
	var categories categories
	db.Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

func getQuestions(c echo.Context) error {
	firstID, err := strconv.Atoi(c.QueryParam("first_id"))
	if err != nil {
		firstID = math.MaxInt64
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	var questions questions
	db.Where("id <= ?", firstID).Order("id desc").Preload("Category").Limit(limit).Find(&questions)
	return c.JSON(http.StatusOK, questions)
}

func createQuestion(c echo.Context) error {
	// TODO: エラーレスポンスをJSONにする
	req := new(questionRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	q := newQuestion(req.Title, req.Body, req.CategoryID)
	db.Create(&q)
	return c.JSON(http.StatusCreated, q)
}

// Request
type questionRequest struct {
	Title  string `json:"title" validate:"required,min=1,max=255"`
	Body string   `json:"body" validate:"required,min=1,max=5000"`
	CategoryID int   `json:"categoryId" validate:"required,categoryId"`
}
type questionsRequest struct {
	FirstID int `query:"first_id" validate:"required"`
	Limit   int `query:"limit" validate:"required"`
}

// Model
type category struct {
	ID   int    `json:"id" gorm:"column:id;primary_key"`
	Name string `json:"name" gorm:"column:name"`
}

type categories = []category

type question struct {
	ID int   `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Title string   `json:"title" gorm:"column:title"`
	Body string   `json:"body" gorm:"column:body"`
	CategoryID int   `json:"categoryId" gorm:"column:category_id"`
	Category category   `json:"category"`
}

type questions = []question

func newQuestion(title string, body string, categoryId int) question {
	return question{Title: title, Body: body, CategoryID: categoryId}
}
