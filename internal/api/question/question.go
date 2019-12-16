package question

import (
	model2 "github.com/c8112002/bbs_api/internal/api/model"
)

type Store interface {
	List(sinceID int, limit int) (model2.Questions, error)
	TotalCount() (int, error)
	CreateQuestion(q *model2.Question) error
}
