package routes

import (
	"github.com/tkdnbb/bookofben-api/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// SetupRoutes configures and returns the router
func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	// CORS configuration
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Initialize handlers
	bibleHandler := handlers.NewBibleHandler()

	// Bible passage routes
	r.Get("/{reference}", bibleHandler.GetBiblePassage)

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/translations", bibleHandler.GetTranslations)
		r.Get("/books", bibleHandler.GetBooks)
	})

	return r
}
