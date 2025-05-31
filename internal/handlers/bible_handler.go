package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/tkdnbb/bookofben-api/internal/models"
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

// AddVerse handles POST /api/verses
func (h *BibleHandler) AddVerse(w http.ResponseWriter, r *http.Request) {
	var verse models.Verse
	if err := json.NewDecoder(r.Body).Decode(&verse); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// 验证必填字段
	if verse.BookID == "" || verse.Text == "" || verse.Chapter <= 0 || verse.Verse <= 0 {
		http.Error(w, "Missing required fields: book_id, text, chapter, verse", http.StatusBadRequest)
		return
	}

	err := h.service.AddVerse(verse)
	if err != nil {
		http.Error(w, "Failed to insert verse", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// SearchVerses handles GET /api/search?q=keyword
func (h *BibleHandler) SearchVerses(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	verses, err := h.service.SearchVerses(query)
	if err != nil {
		http.Error(w, "Search failed", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"query":   query,
		"count":   len(verses),
		"results": verses,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
