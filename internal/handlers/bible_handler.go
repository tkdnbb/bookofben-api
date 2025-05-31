package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/tkdnbb/bookofben-api/internal/services"
)

// BibleHandler handles HTTP requests for Bible API
type BibleHandler struct {
	service *services.BibleService
}

// NewBibleHandler creates a new BibleHandler instance
func NewBibleHandler() *BibleHandler {
	return &BibleHandler{
		service: services.NewBibleService(),
	}
}

// GetBiblePassage handles GET /{reference}
func (h *BibleHandler) GetBiblePassage(w http.ResponseWriter, r *http.Request) {
	// Decode URL parameter
	reference, _ := url.QueryUnescape(chi.URLParam(r, "reference"))
	translation := r.URL.Query().Get("translation")

	response, err := h.service.GetPassage(reference, translation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

// GetTranslations handles GET /api/translations
func (h *BibleHandler) GetTranslations(w http.ResponseWriter, r *http.Request) {
	translations := h.service.GetTranslations()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(translations)
}

// GetBooks handles GET /api/books
func (h *BibleHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books := h.service.GetBooks()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(books)
}
