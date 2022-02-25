package errors

type BaseError interface {
	Error() map[string]interface{}
	HttpStatusCode() int
}