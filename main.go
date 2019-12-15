package main

import (
	"github.com/c8112002/bbs_api/app/db"
	"github.com/c8112002/bbs_api/app/handler"
	"github.com/c8112002/bbs_api/app/router"
	"github.com/c8112002/bbs_api/app/store"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")

	d := db.New(r)

	cs := store.NewCategoryStore(d)
	qs := store.NewQuestionStore(d)
	h := handler.NewHandler(cs, qs)
	h.Register(v1)

	r.Logger.Fatal(r.Start(":1234"))
}
