package libraryhttp

import (
	"encoding/json"
	"errors"
	"homework/library"
	"net/http"
	"time"
)

func statusFromError(err error) int {
	switch {
	case errors.As(err, &ValidationError{}):
		return http.StatusBadRequest
	case errors.Is(err, library.ErrBookNotFound):
		return http.StatusNotFound
	case errors.Is(err, library.ErrBookAlreadyExists):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	errDto := ErrorDto{
		Message: err.Error(),
		Time:    time.Now(),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(errDto); err != nil {
		panic(err)
	}
}

func WriteAppError(w http.ResponseWriter, err error) {
	WriteError(w, statusFromError(err), err)
}

func WriteJSON(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(statusCode)

	if _, err := w.Write(b); err != nil {
		panic(err)
	}
}
