package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository handles database operations
type Repository struct {
	db *mongo.Database
}

// NewRepository creates a new repository instance
func NewRepository() *Repository {
	return &Repository{
		db: GetDatabase(),
	}
}

// GetTranslation retrieves a translation by ID
func (r *Repository) GetTranslation(translationID string) (*Translation, error) {
	ctx := context.Background()
	collection := r.db.Collection("translations")

	var translation Translation
	err := collection.FindOne(ctx, bson.M{"_id": translationID}).Decode(&translation)
	if err != nil {
		return nil, fmt.Errorf("translation not found: %w", err)
	}

	return &translation, nil
}

// GetBook retrieves a book by ID
func (r *Repository) GetBook(bookID string) (*Book, error) {
	ctx := context.Background()
	collection := r.db.Collection("books")

	var book Book
	err := collection.FindOne(ctx, bson.M{"_id": bookID}).Decode(&book)
	if err != nil {
		return nil, fmt.Errorf("book not found: %w", err)
	}

	return &book, nil
}

// GetVerses retrieves verses by book, chapter, and optionally verse number
func (r *Repository) GetVerses(bookID string, chapter, verse int) ([]Verse, error) {
	ctx := context.Background()
	collection := r.db.Collection("verses")

	filter := bson.M{"book_id": bookID, "chapter": chapter}
	if verse > 0 {
		filter["verse"] = verse
	}

	findOptions := options.Find()
	findOptions.Sort = bson.M{"verse": 1}
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to find verses: %w", err)
	}
	defer cursor.Close(ctx)

	var verses []Verse
	if err = cursor.All(ctx, &verses); err != nil {
		return nil, fmt.Errorf("failed to decode verses: %w", err)
	}

	return verses, nil
}

// GetAllTranslations retrieves all translations
func (r *Repository) GetAllTranslations() ([]Translation, error) {
	ctx := context.Background()
	collection := r.db.Collection("translations")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch translations: %w", err)
	}
	defer cursor.Close(ctx)

	var translations []Translation
	if err = cursor.All(ctx, &translations); err != nil {
		return nil, fmt.Errorf("failed to decode translations: %w", err)
	}

	return translations, nil
}

// GetAllBooks retrieves all books
func (r *Repository) GetAllBooks() ([]Book, error) {
	ctx := context.Background()
	collection := r.db.Collection("books")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch books: %w", err)
	}
	defer cursor.Close(ctx)

	var books []Book
	if err = cursor.All(ctx, &books); err != nil {
		return nil, fmt.Errorf("failed to decode books: %w", err)
	}

	return books, nil
}

// InsertVerse inserts a new verse
func (r *Repository) InsertVerse(verse Verse) error {
	ctx := context.Background()
	collection := r.db.Collection("verses")

	_, err := collection.InsertOne(ctx, verse)
	if err != nil {
		return fmt.Errorf("failed to insert verse: %w", err)
	}

	return nil
}

// SearchVerses searches for verses containing the query string
func (r *Repository) SearchVerses(query string) ([]Verse, error) {
	ctx := context.Background()
	collection := r.db.Collection("verses")

	filter := bson.M{"text": bson.M{"$regex": query, "$options": "i"}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}
	defer cursor.Close(ctx)

	var verses []Verse
	if err = cursor.All(ctx, &verses); err != nil {
		return nil, fmt.Errorf("failed to decode search results: %w", err)
	}

	return verses, nil
}
