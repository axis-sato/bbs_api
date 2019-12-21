package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Error struct {
	Error   string
	Details []errorDetail
}

type errorDetail struct {
	Code  string
	Field string
}

func NewError(err error) Error {
	switch e := err.(type) {
	case *echo.HTTPError:
		return Error{
			Error: e.Message.(string),
		}
	case validator.ValidationErrors:
		return newValidationError(e)
	default:
		return newInternalServerError()
	}
}

func newInternalServerError() Error {
	return Error{Error: "Internal Server Error"}
}

func newValidationError(errs validator.ValidationErrors) Error {
	var e Error
	e.Error = "Validation Error"

	for _, v := range errs {
		ed := errorDetail{
			Code:  v.Tag(),
			Field: v.Field(),
		}
		e.Details = append(e.Details, ed)
	}

	return e
}
