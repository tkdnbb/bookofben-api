package routes

import (
	"log"

	"github.com/tkdnbb/bookofben-api/internal/database"
	"github.com/tkdnbb/bookofben-api/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// SetupRoutes configures and returns the router
func SetupRoutes() *chi.Mux {
	// Initialize database connection
	if err := database.InitMongoDB("mongodb://root:example@localhost:27017", "bible_api"); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Initialize sample data
	if err := database.InitializeData(); err != nil {
		log.Printf("Warning: Failed to initialize data: %v", err)
	}

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
		r.Post("/verses", bibleHandler.AddVerse)    // 新增经文
		r.Get("/search", bibleHandler.SearchVerses) // 搜索经文
	})

	return r
}

// CloseDatabase provides a way to close the database connection
func CloseDatabase() error {
	return database.Close()
}
