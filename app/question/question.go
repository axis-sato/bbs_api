package question

import "github.com/c8112002/bbs_api/app/model"

type Store interface {
	List(sinceID int, limit int) (model.Questions, error)
	TotalCount() (int, error)
	CreateQuestion(q *model.Question) error
}
