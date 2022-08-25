package results

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != "Bearer auth" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(chi.URLParam(r, "id")))
}
