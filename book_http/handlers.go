package book_http

import (
	"encoding/json"
	"homework/book"
	"net/http"
	"strconv"
)

type Handlers struct {
	service book.Service
}

func NewHandlers(service book.Service) *Handlers {
	return &Handlers{service}
}

func (h *Handlers) HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	var bookDto book.CreateBookDto

	if err := json.NewDecoder(r.Body).Decode(&bookDto); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	createdBook, err := h.service.Create(r.Context(), bookDto)

	if err != nil {
		WriteAppError(w, err)
		return
	}

	WriteJSON(w, createdBook, http.StatusCreated)
}

func (h *Handlers) HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	var bookDto book.UpdateBookDto
	if err := json.NewDecoder(r.Body).Decode(&bookDto); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	updatedBook, err := h.service.Update(r.Context(), id, bookDto)

	if err != nil {
		WriteAppError(w, err)
		return
	}

	WriteJSON(w, updatedBook, http.StatusOK)
}

func (h *Handlers) HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		WriteAppError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) HandleGetBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	b, err := h.service.GetByID(r.Context(), id)

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

	books, err := h.service.List(r.Context(), book.BookFilter{
		Name:   name,
		Author: author,
		Read:   read,
	})

	if err != nil {
		WriteAppError(w, err)
		return
	}

	WriteJSON(w, books, http.StatusOK)
}
