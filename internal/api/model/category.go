package model

// Model
type Category struct {
	ID   int    `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`
}

type Categories = []Category
