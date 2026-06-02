package book

import "errors"

var ErrBookNotFound = errors.New("book not found")
var ErrBookAlreadyExists = errors.New("book already exists")

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}
