package database

import "github.com/jinzhu/gorm"

type Database struct {
	DB *gorm.DB
}

func NewDatabase() *Database {
	db, err := gorm.Open("mysql", "bbs:bbspassword@tcp(localhost:3306)/bbs?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		print("error")
	}

	return &Database{DB: db}
}

func (db *Database) Find(out interface{}, where ...interface{}) *gorm.DB {
	return db.DB.Find(out, where...)
}