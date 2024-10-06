package errors

import "net/http"

func NewError(err error) *Error {
	return &Error{
		Message: err.Error(),
		Type:    "INTERNAL_SERVER_ERROR",
		Status:  500,
	}
}

func NewErrorNotFound(message string) *Error {
	return &Error{
		Message: message,
		Type:    "NOT_FOUND",
		Status:  http.StatusNotFound,
	}
}

func NewErrorUnauthorized(message string) *Error {
	return &Error{
		Message: message,
		Type:    "UNAUTHORIZED",
		Status:  http.StatusUnauthorized,
	}
}

func NewErrorBadRequest(message string) *Error {
	return &Error{
		Message: message,
		Type:    "BAD_REQUEST",
		Status:  http.StatusBadRequest,
	}
}
