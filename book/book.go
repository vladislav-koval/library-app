package book

import "time"

type BookUpdate struct {
	Read bool
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
