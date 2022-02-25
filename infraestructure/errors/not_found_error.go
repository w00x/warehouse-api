package errors

import "net/http"

type NotFoundError struct {
	Message string
	Description string
	Code int
	httpStatusCode int
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{message, "NOT_FOUND", 1000, http.StatusNotFound }
}

func (e *NotFoundError) Error() map[string]interface{} {
	err := make(map[string]interface{})
	err["message"] = e.Message
	err["description"] = e.Description
	err["code"] = e.Code
	return err
}

func (e *NotFoundError) HttpStatusCode() int {
	return e.httpStatusCode
}