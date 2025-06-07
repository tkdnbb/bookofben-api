package routes

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tkdnbb/bookofben-api/internal/database"
	"github.com/tkdnbb/bookofben-api/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// SetupRoutes configures and returns the router
func SetupRoutes() *chi.Mux {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found or failed to load")
	}

	// Get MongoDB connection string from environment
	mongoConn := os.Getenv("MONGO_CONNECTION")
	if mongoConn == "" {
		log.Fatal("MONGO_CONNECTION not set in environment")
	}

	// Initialize database connection
	if err := database.InitMongoDB(mongoConn, "bible_api"); err != nil {
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
