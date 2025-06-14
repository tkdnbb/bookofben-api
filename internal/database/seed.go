package database

import (
	"context"
	"fmt"

	"github.com/tkdnbb/bookofben-api/internal/data"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// InitializeData seeds the database with initial data
func InitializeData() error {
	db := GetDatabase()
	ctx := context.Background()

	// Initialize translations
	if err := initializeTranslations(db, ctx); err != nil {
		return fmt.Errorf("failed to initialize translations: %w", err)
	}

	// Initialize books
	if err := initializeBooks(db, ctx); err != nil {
		return fmt.Errorf("failed to initialize books: %w", err)
	}

	// Initialize verses
	if err := initializeVerses(db, ctx); err != nil {
		return fmt.Errorf("failed to initialize verses: %w", err)
	}

	return nil
}

func initializeTranslations(db *mongo.Database, ctx context.Context) error {
	collection := db.Collection("translations")

	count, _ := collection.CountDocuments(ctx, bson.M{})
	if count > 0 {
		return nil // Already initialized
	}

	translations := []interface{}{
		Translation{ID: "cuv", Name: "Chinese Union Version", Note: "Public Domain"},
		Translation{ID: "kjv", Name: "King James Version", Note: "Public Domain"},
		Translation{ID: "en", Name: "English Version", Note: "Public Domain"},
	}

	_, err := collection.InsertMany(ctx, translations)
	if err != nil {
		return err
	}

	fmt.Println("Initialized translations")
	return nil
}

func initializeBooks(db *mongo.Database, ctx context.Context) error {
	collection := db.Collection("books")

	count, _ := collection.CountDocuments(ctx, bson.M{})
	if count > 0 {
		return nil // Already initialized
	}

	books := []interface{}{
		Book{ID: "GEN", Name: "創世紀", Chapters: 50},
		Book{ID: "MAT", Name: "馬太福音", Chapters: 28},
		Book{ID: "JHN", Name: "約翰福音", Chapters: 21},
		Book{ID: "BEN", Name: "The Book of Jachanan Ben Kathryn", Chapters: 73},
	}

	_, err := collection.InsertMany(ctx, books)
	if err != nil {
		return err
	}

	fmt.Println("Initialized books")
	return nil
}

func initializeVerses(db *mongo.Database, ctx context.Context) error {
	collection := db.Collection("verses")

	count, _ := collection.CountDocuments(ctx, bson.M{})
	if count > 0 {
		return nil // Already initialized
	}

	var verses []any

	// Genesis Chapter 1 (KJV) - 保持原有数据
	verses = append(verses,
		Verse{BookID: "GEN", TranslationID: "kjv", BookName: "Genesis", Chapter: 1, Verse: 1, Text: "In the beginning God created the heaven and the earth."},
		Verse{BookID: "GEN", TranslationID: "kjv", BookName: "Genesis", Chapter: 1, Verse: 2, Text: "And the earth was without form, and void; and darkness was upon the face of the deep. And the Spirit of God moved upon the face of the waters."},
		Verse{BookID: "GEN", TranslationID: "kjv", BookName: "Genesis", Chapter: 1, Verse: 3, Text: "And God said, Let there be light: and there was light."},
		Verse{BookID: "GEN", TranslationID: "kjv", BookName: "Genesis", Chapter: 1, Verse: 4, Text: "And God saw the light, that it was good: and God divided the light from the darkness."},
		Verse{BookID: "GEN", TranslationID: "kjv", BookName: "Genesis", Chapter: 1, Verse: 5, Text: "And God called the light Day, and the darkness he called Night. And the evening and the morning were the first day."},
	)

	// John 3:16 (Chinese) - 保持原有数据
	verses = append(verses,
		Verse{BookID: "JHN", TranslationID: "cuv", BookName: "約翰福音", Chapter: 3, Verse: 16, Text: "神愛世人，甚至將他的獨生子賜給他們，叫一切信他的，不至滅亡，反得永生。"},
	)

	// The Book of Jachanan Ben Kathryn - 批量加载章节
	benChapters := data.GetTotalChapters()
	totalBenVerses := 0

	for chapterNum := 1; chapterNum <= benChapters; chapterNum++ {
		chapterVerses := data.GetChapterVerses(chapterNum)
		for i, verseText := range chapterVerses {
			verses = append(verses, Verse{
				BookID:        "BEN",
				BookName:      "The Book of Jachanan Ben Kathryn",
				TranslationID: "en",
				Chapter:       chapterNum,
				Verse:         i + 1,
				Text:          verseText,
			})
		}
		totalBenVerses += len(chapterVerses)
		fmt.Printf("- Book of Ben Chapter %d: %d verses\n", chapterNum, len(chapterVerses))
	}

	_, err := collection.InsertMany(ctx, verses)
	if err != nil {
		return err
	}

	fmt.Printf("Initialized verses: %d total\n", len(verses))
	fmt.Printf("- Book of Ben total: %d verses\n", totalBenVerses)
	return nil
}
