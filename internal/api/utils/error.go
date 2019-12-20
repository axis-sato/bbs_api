package utils

type Error struct {
	Message string
	Errors  []errorDetail
}

type errorDetail struct {
	Code  string
	Field string
}

func NewInternalServerError() Error {
	return Error{Message: "Internal Server Error"}
}
