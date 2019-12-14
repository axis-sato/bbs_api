package db

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/jinzhu/gorm"
)

func New(r *echo.Echo) *gorm.DB {
	db, err := gorm.Open("mysql", "bbs:bbspassword@tcp(localhost:3306)/bbs?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	db.SetLogger(r.Logger)
	return db
}
