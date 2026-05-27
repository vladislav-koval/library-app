package book

import "time"

type Book struct {
	Name       string
	Author     string
	TotalPages int
	Read       bool

	CreatedAt time.Time
	ReadAt    *time.Time
}

type BookFilter struct {
	Name   string
	Author string
	Read   *bool
}

type BookUpdate struct {
	Read bool
}

func NewBook(name string, author string, totalPages int) Book {
	return Book{
		Name:       name,
		Author:     author,
		TotalPages: totalPages,
		Read:       false,

		CreatedAt: time.Now(),
		ReadAt:    nil,
	}
}

func (b *Book) ReadBook() {
	completeTime := time.Now()

	b.Read = true
	b.ReadAt = &completeTime
}

func (b *Book) UnreadBook() {
	b.Read = false
	b.ReadAt = nil
}
