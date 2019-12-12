package eintities

type Category struct {
	ID int
	Name string
}

func NewCategory(id int, name string) *Category {
	return &Category{
		ID:   id,
		Name: name,
	}
}