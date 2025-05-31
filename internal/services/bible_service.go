package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tkdnbb/bookofben-api/internal/data"
	"github.com/tkdnbb/bookofben-api/internal/models"
)

// BibleService handles business logic for Bible operations
type BibleService struct {
	data *data.BibleData
}

// NewBibleService creates a new BibleService instance
func NewBibleService() *BibleService {
	return &BibleService{
		data: data.GetBibleData(),
	}
}

// GetPassage retrieves a Bible passage by reference and translation
func (s *BibleService) GetPassage(reference, translation string) (*models.BibleResponse, error) {
	// Default translation
	if translation == "" {
		translation = "cuv"
	}

	// Check if translation exists
	trans, exists := s.data.Translations[translation]
	if !exists {
		return nil, fmt.Errorf("translation not found")
	}

	// Parse reference
	bookKey, chapterKey := s.parseReference(reference)
	if bookKey == "" {
		return nil, fmt.Errorf("invalid reference format")
	}

	// Check if book exists
	_, exists = s.data.Books[bookKey]
	if !exists {
		return nil, fmt.Errorf("book not found")
	}

	// Get verses
	verses := s.getVerses(translation, chapterKey)
	if len(verses) == 0 {
		return nil, fmt.Errorf("chapter not found")
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

// GetTranslations returns all available translations
func (s *BibleService) GetTranslations() []models.Translation {
	translations := make([]models.Translation, 0, len(s.data.Translations))
	for _, translation := range s.data.Translations {
		translations = append(translations, translation)
	}
	return translations
}

// GetBooks returns all available books
func (s *BibleService) GetBooks() []models.Book {
	books := make([]models.Book, 0, len(s.data.Books))
	for _, book := range s.data.Books {
		books = append(books, book)
	}
	return books
}

func (s *BibleService) parseReference(reference string) (string, string) {
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
		return "", ""
	}

	// Join all parts except the last as the book name
	bookName := strings.ToLower(strings.Join(parts[:len(parts)-1], " "))
	bookID, exists := bookMap[bookName]
	if !exists {
		// Try direct matching as before
		for name, id := range bookMap {
			if strings.Contains(strings.ToLower(reference), strings.ToLower(name)) {
				bookID = id
				break
			}
		}
		if bookID == "" {
			return "", ""
		}
	}

	chapterPart := parts[len(parts)-1]
	if strings.Contains(chapterPart, ":") {
		// Contains verse number (e.g., "3:16")
		chapterVerse := strings.Split(chapterPart, ":")
		chapter, _ := strconv.Atoi(chapterVerse[0])
		verse, _ := strconv.Atoi(chapterVerse[1])
		return bookID, fmt.Sprintf("%s.%d.%d", bookID, chapter, verse)
	} else {
		// Only chapter number (e.g., "1")
		chapter, _ := strconv.Atoi(chapterPart)
		return bookID, fmt.Sprintf("%s.%d", bookID, chapter)
	}
}

func (s *BibleService) getVerses(translation, key string) []models.Verse {
	if translationData, exists := s.data.Verses[translation]; exists {
		if verses, exists := translationData[key]; exists {
			return verses
		}
	}
	return nil
}

func (s *BibleService) buildText(verses []models.Verse) string {
	var textParts []string
	for _, verse := range verses {
		textParts = append(textParts, verse.Text)
	}
	return strings.Join(textParts, "")
}
