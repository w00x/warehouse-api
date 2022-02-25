package errors

import "net/http"

type InternalServerError struct {
	Message string
	Description string
	Code int
	httpStatusCode int
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{message, "INTERNAL_SERVER_ERROR", 1001, http.StatusInternalServerError }
}

func(e *InternalServerError) Error() map[string]interface{} {
	err := make(map[string]interface{})
	err["message"] = e.Message
	err["description"] = e.Description
	err["code"] = e.Code
	return err
}

func (e *InternalServerError) HttpStatusCode() int {
	return e.httpStatusCode
}