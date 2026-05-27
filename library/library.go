package library

import (
	"homework/book"
	"sync"
)

type Library struct {
	books map[string]book.Book
	mtx   sync.RWMutex
}

func NewLibrary() *Library {
	return &Library{
		books: make(map[string]book.Book),
		mtx:   sync.RWMutex{},
	}
}

func (l *Library) AddBook(b book.Book) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	newBook := book.NewBook(b.Name, b.Author, b.TotalPages)

	if _, ok := l.books[b.Name]; ok {
		return ErrBookAlreadyExists
	}

	l.books[b.Name] = newBook

	return nil
}

func (l *Library) GetBook(name string) (book.Book, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	b, ok := l.books[name]

	if !ok {
		return book.Book{}, ErrBookNotFound
	}

	return b, nil
}

func (l *Library) RemoveBook(name string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.books[name]; !ok {
		return ErrBookNotFound
	}

	delete(l.books, name)

	return nil
}

func (l *Library) GetBooksByFilter(filters book.BookFilter) []book.Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	books := make([]book.Book, 0)

	for _, b := range l.books {
		if filters.Name != "" && filters.Name != b.Name {
			continue
		}

		if filters.Author != "" && filters.Author != b.Author {
			continue
		}

		if filters.Read != nil && *filters.Read != b.Read {
			continue
		}

		books = append(books, b)
	}

	return books
}

func (l *Library) UpdateBook(name string, data book.BookUpdate) (book.Book, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	b, ok := l.books[name]

	if !ok {
		return book.Book{}, ErrBookNotFound
	}

	if data.Read {
		b.ReadBook()
	} else {
		b.UnreadBook()
	}

	l.books[name] = b

	return b, nil
}
