package errors

import "os"

type Error struct {
	Message string
	Status  int
	Type    string
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) SetMessage(message string) *Error {
	e.Message = message
	return e
}

func (e *Error) SetStatus(status int) *Error {
	e.Status = status
	return e
}

func (e *Error) BadRequest() *Error {
	e.Type = "BAD_REQUEST"
	e.Status = 400
	return e
}

func (e *Error) Unauthorized() *Error {
	e.Status = 401
	return e
}

func (e *Error) Forbidden() *Error {
	e.Status = 403
	return e
}

func (e *Error) NotFound() *Error {
	e.Status = 404
	return e
}

func (e *Error) SetTypeSystem() *Error {
	e.Type = "system"
	os.Exit(1)
	return e
}
