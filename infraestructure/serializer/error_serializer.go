package serializer

import (
	"warehouse/infraestructure/errors"
)

type ErrorSerializer struct {
	Message string 		`json:"message"`
	Description string 	`json:"description"`
	Code int 			`json:"code"`
}

func NewErrorSerializerFromError(error *errors.NotFoundError) *ErrorSerializer {
	return &ErrorSerializer{error.Message, error.Description, error.Code}
}