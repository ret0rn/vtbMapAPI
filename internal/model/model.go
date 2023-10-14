package model

import "fmt"

type Success struct {
	Message string `json:"message"`
}

type Error struct {
	Error string `json:"error"`
}

func NewSuccess(format string, args ...interface{}) Success {
	return Success{Message: fmt.Sprintf(format, args...)}
}

func NewError(format string, args ...interface{}) Error {
	return Error{Error: fmt.Sprintf(format, args...)}
}
