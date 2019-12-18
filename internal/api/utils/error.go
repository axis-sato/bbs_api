package utils

type Error struct {
	Code  string
	Field string
}

func NewInternalServerError() Error {
	return Error{Code: "internal_server_error"}
}
