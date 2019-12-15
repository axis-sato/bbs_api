package model

import "time"

type Question struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Title      string    `gorm:"column:title"`
	Body       string    `gorm:"column:body"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	CategoryID int       `gorm:"column:category_id"`
	Category   Category
}

type Questions = []Question
