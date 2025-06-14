package database

// Verse represents a Bible verse in the database
type Verse struct {
	BookID        string `json:"book_id" bson:"book_id"`
	TranslationID string `json:"translation_id" bson:"tranlation_id"`
	BookName      string `json:"book_name" bson:"book_name"`
	Chapter       int    `json:"chapter" bson:"chapter"`
	Verse         int    `json:"verse" bson:"verse"`
	Text          string `json:"text" bson:"text"`
}

// Translation represents a Bible translation
type Translation struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Note string `json:"note" bson:"note"`
}

// Book represents a Bible book
type Book struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Chapters int    `json:"chapters" bson:"chapters"`
}
