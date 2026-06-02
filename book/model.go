package book

import "time"

type Book struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Author     string `json:"author"`
	TotalPages int    `json:"totalPages"`
	Read       bool   `json:"read"`

	CreatedAt time.Time  `json:"createdAt"`
	ReadAt    *time.Time `json:"readAt"`
}

type BookFilter struct {
	Name   string
	Author string
	Read   *bool
}
