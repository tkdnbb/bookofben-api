package models

// Verse represents a single Bible verse
type Verse struct {
	BookID   string `json:"book_id"`
	BookName string `json:"book_name"`
	Chapter  int    `json:"chapter"`
	Verse    int    `json:"verse"`
	Text     string `json:"text"`
}

// BibleResponse represents the API response for Bible passages
type BibleResponse struct {
	Reference       string  `json:"reference"`
	Verses          []Verse `json:"verses"`
	Text            string  `json:"text"`
	TranslationID   string  `json:"translation_id"`
	TranslationName string  `json:"translation_name"`
	TranslationNote string  `json:"translation_note"`
}

// Translation represents a Bible translation
type Translation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Note string `json:"note"`
}

// Book represents a Bible book
type Book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Chapters int    `json:"chapters"`
}
