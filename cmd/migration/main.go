package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	migration "github.com/rubenv/sql-migrate"
)

func main() {
	db, err := gorm.Open("mysql", "bbs:bbspassword@tcp(localhost:3306)/bbs?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)

	flag.Parse()
	cmd := flag.Arg(0)

	if cmd == "migrate" {
		migrate(db)
	} else if cmd == "seed" {
		seed(db)
	} else {
		fmt.Println("実行可能なコマンドは以下です。")
		fmt.Println("go run database/main.go [migrate|seed]")
	}
}

func migrate(db *gorm.DB) {

	migrations := &migration.FileMigrationSource{
		Dir: "configs/migration/migrations",
	}

	cmd := flag.Arg(1)

	m := func(dir migration.MigrationDirection) {
		n, err := migration.Exec(db.DB(), "mysql", migrations, dir)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("Applied %d migrations!\n", n)
	}

	if cmd == "up" {
		m(migration.Up)
	} else if cmd == "down" {
		m(migration.Down)
	} else if cmd == "refresh" {
		m(migration.Down)
		m(migration.Up)
	} else {
		fmt.Println("実行可能なコマンドは以下です。")
		fmt.Println("go run cmd/migration/main.go migrate [up|down]")
	}
}

func seed(db *gorm.DB) {
	fmt.Println("run seeding")

	dir := "configs/migration/seeds"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err.Error())
	}

	var paths []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	for _, path := range paths {
		executeSQL(db, path)
	}
}

func executeSQL(db *gorm.DB, file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	b, err := ioutil.ReadAll(f)
	sql := string(b)

	db.Exec(sql)
}
