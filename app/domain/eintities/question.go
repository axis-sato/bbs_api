package eintities

import "time"

type Question struct {
	ID int
	Title string
	Body string
	CreatedAt time.Time
	Category Category
}

func NewQuestion(id int, title string, body string, category Category) *Question {
	return &Question{
		ID:        id,
		Title:     title,
		Body:      body,
		CreatedAt: time.Now(),
		Category:  category,
	}
}