package db

import (
	"fmt"

	migration "github.com/rubenv/sql-migrate"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
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

func TestDB() *gorm.DB {
	db, err := gorm.Open("mysql", "bbs:bbspassword@tcp(localhost:3307)/bbs?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

func DropTestDB() error {
	db, err := gorm.Open("mysql", "bbs:bbspassword@tcp(localhost:3307)/bbs?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	migrations := &migration.FileMigrationSource{
		Dir: "../../../configs/migration/migrations",
	}
	n, err := migration.Exec(db.DB(), "mysql", migrations, migration.Down)
	if err != nil {
		return err
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return nil
}

func Migrate(db *gorm.DB) error {
	migrations := &migration.FileMigrationSource{
		Dir: "../../../configs/migration/migrations",
	}
	n, err := migration.Exec(db.DB(), "mysql", migrations, migration.Up)
	if err != nil {
		return err
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return nil
}
