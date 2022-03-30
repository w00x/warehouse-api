package dto

import (
	"warehouse/infraestructure/errors"
)

type ErrorDto struct {
	Message string 		`json:"message"`
	Description string 	`json:"description"`
	Code int 			`json:"code"`
}

func NewErrorDtoFromError(error *errors.NotFoundError) *ErrorDto {
	return &ErrorDto{error.Message, error.Description, error.Code}
}