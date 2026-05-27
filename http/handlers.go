package libraryhttp

import (
	"encoding/json"
	"homework/book"
	"homework/library"
	"net/http"
	"strconv"
)

type Handlers struct {
	library *library.Library
}

func NewHandlers(library *library.Library) *Handlers {
	return &Handlers{library}
}

func (h *Handlers) HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	var bookDto BookDto

	if err := json.NewDecoder(r.Body).Decode(&bookDto); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := bookDto.ValidateForCreate(); err != nil {
		WriteAppError(w, err)
		return
	}

	b := book.Book{
		Name:       bookDto.Name,
		Author:     bookDto.Author,
		TotalPages: bookDto.TotalPages,
	}

	if err := h.library.AddBook(b); err != nil {
		WriteAppError(w, err)
		return
	}

	WriteJSON(w, bookDto, http.StatusCreated)
}

func (h *Handlers) HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	var bookDto UpdateBookDto
	if err := json.NewDecoder(r.Body).Decode(&bookDto); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := bookDto.Validate(); err != nil {
		WriteAppError(w, err)
		return
	}

	bookUpdate := book.BookUpdate{
		Read: *bookDto.Read,
	}

	b, err := h.library.UpdateBook(name, bookUpdate)

	if err != nil {
		WriteAppError(w, err)
		return
	}

	WriteJSON(w, b, http.StatusOK)
}

func (h *Handlers) HandleGetBook(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	b, err := h.library.GetBook(name)

	if err != nil {
		WriteAppError(w, err)
		return
	}

	WriteJSON(w, b, http.StatusOK)
}

func (h *Handlers) HandleFilterBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	name := query.Get("name")
	author := query.Get("author")
	readStr := query.Get("read")
	var read *bool

	if readStr != "" {
		r, err := strconv.ParseBool(readStr)

		if err != nil {
			WriteError(w, http.StatusBadRequest, err)
			return
		}
		read = &r
	}

	books := h.library.GetBooksByFilter(book.BookFilter{
		Name:   name,
		Author: author,
		Read:   read,
	})

	WriteJSON(w, books, http.StatusOK)
}

func (h *Handlers) HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	if err := h.library.RemoveBook(name); err != nil {
		WriteAppError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
