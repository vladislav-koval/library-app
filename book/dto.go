package book

import (
	"encoding/json"
	"time"
)

type CreateBookDto struct {
	Name       string `json:"name"`
	Author     string `json:"author"`
	TotalPages int    `json:"totalPages"`
}

func (b CreateBookDto) ValidateForCreate() error {
	if b.Name == "" {
		return ValidationError{"name is required"}
	}
	if b.Author == "" {
		return ValidationError{"author is required"}
	}
	if b.TotalPages <= 0 {
		return ValidationError{"totalPages should be >= 0"}
	}

	return nil
}

type UpdateBookDto struct {
	Name       *string `json:"name"`
	Author     *string `json:"author"`
	TotalPages *int    `json:"totalPages"`
	Read       *bool   `json:"read"`
}

func (b UpdateBookDto) ValidateForUpdate() error {
	if b.Name == nil &&
		b.Author == nil &&
		b.TotalPages == nil &&
		b.Read == nil {
		return ValidationError{"empty update payload"}
	}

	if b.Name != nil && *b.Name == "" {
		return ValidationError{"name should not be empty"}
	}

	if b.Author != nil && *b.Author == "" {
		return ValidationError{"author should not be empty"}
	}

	if b.TotalPages != nil && *b.TotalPages <= 0 {
		return ValidationError{"totalPages should be > 0"}
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
