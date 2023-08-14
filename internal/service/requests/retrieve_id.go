package requests

import (
	"errors"
	"github.com/go-chi/chi"
	"net/http"
)

func retrieveId(r *http.Request) (*string, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, errors.New("")
	}
	return &id, nil
}
