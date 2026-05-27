package libraryhttp

import (
	"encoding/json"
	"time"
)

type BookDto struct {
	Name       string `json:"name"`
	Author     string `json:"author"`
	TotalPages int    `json:"totalPages"`
}

type UpdateBookDto struct {
	Read *bool `json:"read"`
}

func (b UpdateBookDto) Validate() error {
	if b.Read == nil {
		return ValidationError{Message: "read is required"}
	}

	return nil
}

func (b BookDto) ValidateForCreate() error {
	if b.Name == "" {
		return ValidationError{"name is required"}
	}
	if b.Author == "" {
		return ValidationError{"author is required"}
	}
	if b.TotalPages == 0 {
		return ValidationError{"totalPages is required"}
	}

	if b.TotalPages < 0 {
		return ValidationError{"totalPages should be >= 0"}
	}

	return nil
}

type ErrorDto struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func (e ErrorDto) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return string(b)
}
