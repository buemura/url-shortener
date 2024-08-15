package handler

import (
	"encoding/json"
	"net/http"

	"github.com/a-h/templ"
	"github.com/buemura/url-shortener/views"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	http.Handler
}

func NewReader() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(s *chi.Mux) http.Handler {
	s.Get("/", h.renderIndex)
	s.Post("/short", h.calculateBonus)
	return s
}

func (h *Handler) renderIndex(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.Index()).ServeHTTP(w, r)
}

func (h *Handler) calculateBonus(w http.ResponseWriter, r *http.Request) {
	// Get url input
	url := r.FormValue("url")

	// components.EmployeeSales(validRecords).Render(r.Context(), w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"url": url})
}
