package services

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/tkdnbb/bookofben-api/internal/database"
	"github.com/tkdnbb/bookofben-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BibleService handles business logic for Bible operations
type BibleService struct {
	repo *database.Repository
}

// NewBibleService creates a new BibleService instance
func NewBibleService() *BibleService {
	return &BibleService{
		repo: database.NewRepository(), // 你需要确保这个方法存在
	}
}

// GetPassage retrieves a Bible passage by reference and translation
func (s *BibleService) GetPassage(reference, translation string) (*models.BibleResponse, error) {
	// Default translation
	if translation == "" {
		translation = "en" // 或者你想要的默认翻译
	}

	// Parse reference
	bookID, chapter, startVerse, endVerse, err := s.parseReference(reference)
	if err != nil {
		return nil, err
	}

	// Get translation info from database
	trans, err := s.repo.GetTranslation(translation)
	if err != nil {
		return nil, fmt.Errorf("translation not found")
	}

	// Get verses from database
	verses, err := s.getVersesFromDB(bookID, chapter, startVerse, endVerse)
	if err != nil {
		return nil, err
	}

	if len(verses) == 0 {
		return nil, fmt.Errorf("no verses found")
	}

	// Build response
	response := &models.BibleResponse{
		Reference:       reference,
		Verses:          verses,
		Text:            s.buildText(verses),
		TranslationID:   trans.ID,
		TranslationName: trans.Name,
		TranslationNote: trans.Note,
	}

	return response, nil
}

func (s *BibleService) parseReference(reference string) (bookID string, chapter int, startVerse int, endVerse int, err error) {
	// Book name mapping
	bookMap := map[string]string{
		"創世記":                              "GEN",
		"創世紀":                              "GEN",
		"馬太福音":                             "MAT",
		"約翰福音":                             "JHN",
		"genesis":                          "GEN",
		"matthew":                          "MAT",
		"john":                             "JHN",
		"the book of jachanan ben kathryn": "BEN",
	}

	parts := strings.Fields(reference)
	if len(parts) < 2 {
		return "", 0, 0, 0, fmt.Errorf("invalid reference format")
	}

	// Join all parts except the last as the book name
	bookName := strings.ToLower(strings.Join(parts[:len(parts)-1], " "))
	bookID, exists := bookMap[bookName]
	if !exists {
		// Try direct matching
		for name, id := range bookMap {
			if strings.Contains(strings.ToLower(reference), strings.ToLower(name)) {
				bookID = id
				break
			}
		}
		if bookID == "" {
			return "", 0, 0, 0, fmt.Errorf("book not found")
		}
	}

	chapterPart := parts[len(parts)-1]
	if strings.Contains(chapterPart, ":") {
		// Contains verse number (e.g., "3:16" or "1:1-9")
		chapterVerse := strings.Split(chapterPart, ":")
		chapter, _ = strconv.Atoi(chapterVerse[0])

		versePart := chapterVerse[1]
		if strings.Contains(versePart, "-") {
			// Range of verses (e.g., "1-9")
			verseRange := strings.Split(versePart, "-")
			startVerse, _ = strconv.Atoi(verseRange[0])
			endVerse, _ = strconv.Atoi(verseRange[1])
		} else {
			// Single verse
			startVerse, _ = strconv.Atoi(versePart)
			endVerse = startVerse
		}
	} else {
		// Only chapter number (e.g., "1") - return entire chapter
		chapter, _ = strconv.Atoi(chapterPart)
		startVerse = 0 // 0 means all verses in the chapter
		endVerse = 0
	}

	return bookID, chapter, startVerse, endVerse, nil
}

func (s *BibleService) getVersesFromDB(bookID string, chapter int, startVerse int, endVerse int) ([]models.Verse, error) {
	ctx := context.Background()
	collection := database.GetDatabase().Collection("verses")

	// Build filter
	filter := bson.M{
		"book_id": bookID,
		"chapter": chapter,
	}

	// Add verse filter if specified
	if startVerse > 0 {
		if endVerse > 0 && endVerse != startVerse {
			filter["verse"] = bson.M{"$gte": startVerse, "$lte": endVerse}
		} else {
			filter["verse"] = startVerse
		}
	}

	// Sort by verse number
	opts := options.Find().SetSort(bson.D{{Key: "verse", Value: 1}})

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var dbVerses []database.Verse
	if err = cursor.All(ctx, &dbVerses); err != nil {
		return nil, err
	}

	// Convert to models.Verse
	verses := make([]models.Verse, len(dbVerses))
	for i, dbVerse := range dbVerses {
		verses[i] = models.Verse{
			BookID:   dbVerse.BookID,
			BookName: dbVerse.BookName,
			Chapter:  dbVerse.Chapter,
			Verse:    dbVerse.Verse,
			Text:     dbVerse.Text,
		}
	}

	return verses, nil
}

// GetTranslations returns all available translations
func (s *BibleService) GetTranslations() []models.Translation {
	// 从数据库获取翻译信息
	ctx := context.Background()
	collection := database.GetDatabase().Collection("translations")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []models.Translation{}
	}
	defer cursor.Close(ctx)

	var dbTranslations []database.Translation
	cursor.All(ctx, &dbTranslations)

	translations := make([]models.Translation, len(dbTranslations))
	for i, dbTrans := range dbTranslations {
		translations[i] = models.Translation{
			ID:   dbTrans.ID,
			Name: dbTrans.Name,
			Note: dbTrans.Note,
		}
	}

	return translations
}

// GetBooks returns all available books
func (s *BibleService) GetBooks() []models.Book {
	// 从数据库获取书籍信息
	ctx := context.Background()
	collection := database.GetDatabase().Collection("books")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []models.Book{}
	}
	defer cursor.Close(ctx)

	var dbBooks []database.Book
	cursor.All(ctx, &dbBooks)

	books := make([]models.Book, len(dbBooks))
	for i, dbBook := range dbBooks {
		books[i] = models.Book{
			ID:       dbBook.ID,
			Name:     dbBook.Name,
			Chapters: dbBook.Chapters,
		}
	}

	return books
}

func (s *BibleService) buildText(verses []models.Verse) string {
	var textParts []string
	for _, verse := range verses {
		textParts = append(textParts, verse.Text)
	}
	return strings.Join(textParts, " ") // 用空格连接而不是直接连接
}

// AddVerse adds a new verse to the database
func (s *BibleService) AddVerse(verse models.Verse) error {
	dbVerse := database.Verse{
		BookID:   verse.BookID,
		BookName: verse.BookName,
		Chapter:  verse.Chapter,
		Verse:    verse.Verse,
		Text:     verse.Text,
	}
	return s.repo.InsertVerse(dbVerse)
}

// SearchVerses searches for verses containing the query string
func (s *BibleService) SearchVerses(query string) ([]models.Verse, error) {
	dbVerses, err := s.repo.SearchVerses(query)
	if err != nil {
		return nil, err
	}

	verses := make([]models.Verse, len(dbVerses))
	for i, dbVerse := range dbVerses {
		verses[i] = models.Verse{
			BookID:   dbVerse.BookID,
			BookName: dbVerse.BookName,
			Chapter:  dbVerse.Chapter,
			Verse:    dbVerse.Verse,
			Text:     dbVerse.Text,
		}
	}

	return verses, nil
}
