package handler

import (
	"encoding/json"
	"net/http"

	"github.com/a-h/templ"
	"github.com/buemura/url-shortener/internal/core/usecase"
	"github.com/buemura/url-shortener/internal/infra/cache"
	"github.com/buemura/url-shortener/internal/infra/database"
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
	s.Post("/shorten", h.calculateBonus)
	s.Get("/{code}", h.getUrl)
	return s
}

func (h *Handler) renderIndex(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.Index()).ServeHTTP(w, r)
}

func (h *Handler) getUrl(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	// Call usecase
	db := database.NewPgxUrlRepository()
	redis := cache.NewRedisCacheRepository()
	uc := usecase.NewGetShortenedUrl(redis, db)

	url, err := uc.Execute(code)
	if err != nil {
		HandleRequestError(w, http.StatusInternalServerError, err, "")
		return
	}

	// FIXME: Redirect to original url
	w.WriteHeader(http.StatusFound)
	http.Redirect(w, r, url.OriginalUrl, http.StatusMovedPermanently)
}

func (h *Handler) calculateBonus(w http.ResponseWriter, r *http.Request) {
	// Get url input
	urlInput := r.FormValue("url")

	// Call usecase
	db := database.NewPgxUrlRepository()
	uc := usecase.NewCreateShortenedUrl(db)

	url, err := uc.Execute(urlInput)
	if err != nil {
		HandleRequestError(w, http.StatusInternalServerError, err, "")
		return
	}

	// components.EmployeeSales(validRecords).Render(r.Context(), w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(url)
}
