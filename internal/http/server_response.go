package http

import (
	"backend-bootcamp-assignment-2024/internal/http/api"
	"backend-bootcamp-assignment-2024/internal/repo/repoerr"
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

func (s *server) json(w http.ResponseWriter, r *http.Request, status int, resp interface{}) {
	render.Status(r, status)
	render.JSON(w, r, resp)
}

func (s *server) error(w http.ResponseWriter, r *http.Request, status int, err error) {
	switch {
	case status == http.StatusBadRequest:
		w.WriteHeader(status)
	case errors.Is(err, repoerr.ErrNotFound):
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, errUnauthorized):
		w.WriteHeader(http.StatusUnauthorized)
	default:
		s.internalServerError(w, r, err)
	}
}

const (
	retryAfterInSec = "30"
)

func (s *server) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Retry-After", retryAfterInSec)

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, api.N5xx{Message: err.Error()})
	}
}
