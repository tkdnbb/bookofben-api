package data

import "github.com/tkdnbb/bookofben-api/internal/models"

// BibleData holds all Bible data in memory
type BibleData struct {
	Books        map[string]models.Book
	Verses       map[string]map[string][]models.Verse // [translation][book.chapter] -> verses
	Translations map[string]models.Translation
}

var instance *BibleData

// GetBibleData returns the singleton instance of BibleData
func GetBibleData() *BibleData {
	if instance == nil {
		instance = initializeBibleData()
	}
	return instance
}

func initializeBibleData() *BibleData {
	bibleData := &BibleData{
		Books:        make(map[string]models.Book),
		Verses:       make(map[string]map[string][]models.Verse),
		Translations: make(map[string]models.Translation),
	}

	// Initialize translations
	bibleData.Translations["cuv"] = models.Translation{
		ID:   "cuv",
		Name: "Chinese Union Version",
		Note: "Public Domain",
	}
	bibleData.Translations["kjv"] = models.Translation{
		ID:   "kjv",
		Name: "King James Version",
		Note: "Public Domain",
	}
	bibleData.Translations["en"] = models.Translation{
		ID:   "en",
		Name: "English Version",
		Note: "Public Domain",
	}

	// Initialize books
	bibleData.Books["GEN"] = models.Book{ID: "GEN", Name: "創世紀", Chapters: 50}
	bibleData.Books["MAT"] = models.Book{ID: "MAT", Name: "馬太福音", Chapters: 28}
	bibleData.Books["JHN"] = models.Book{ID: "JHN", Name: "約翰福音", Chapters: 21}
	bibleData.Books["BEN"] = models.Book{ID: "BEN", Name: "The Book of Jachanan Ben Kathryn", Chapters: 73}

	// Initialize verses data
	bibleData.Verses["cuv"] = make(map[string][]models.Verse)
	bibleData.Verses["kjv"] = make(map[string][]models.Verse)
	bibleData.Verses["en"] = make(map[string][]models.Verse)

	loadSampleVerses(bibleData)

	return bibleData
}

func loadSampleVerses(bibleData *BibleData) {
	// Genesis 1 (KJV)
	genesis1KJV := []models.Verse{
		{BookID: "GEN", BookName: "Genesis", Chapter: 1, Verse: 1, Text: "In the beginning God created the heaven and the earth."},
		{BookID: "GEN", BookName: "Genesis", Chapter: 1, Verse: 2, Text: "And the earth was without form, and void; and darkness was upon the face of the deep. And the Spirit of God moved upon the face of the waters."},
		{BookID: "GEN", BookName: "Genesis", Chapter: 1, Verse: 3, Text: "And God said, Let there be light: and there was light."},
		{BookID: "GEN", BookName: "Genesis", Chapter: 1, Verse: 4, Text: "And God saw the light, that it was good: and God divided the light from the darkness."},
		{BookID: "GEN", BookName: "Genesis", Chapter: 1, Verse: 5, Text: "And God called the light Day, and the darkness he called Night. And the evening and the morning were the first day."},
	}
	bibleData.Verses["kjv"]["GEN.1"] = genesis1KJV

	// John 3:16 (CUV)
	john316CUV := []models.Verse{
		{BookID: "JHN", BookName: "約翰福音", Chapter: 3, Verse: 16, Text: "神愛世人，甚至將他的獨生子賜給他們，叫一切信他的，不至滅亡，反得永生。"},
	}
	bibleData.Verses["cuv"]["JHN.3.16"] = john316CUV

	// The Book of Jachanan Ben Kathryn 1:1 (EN)
	ben11EN := []models.Verse{
		{BookID: "BEN", BookName: "The Book of Jachanan Ben Kathryn", Chapter: 1, Verse: 1, Text: "THE burden of the word of the LORD which came unto John the son of Kathryn, the daughter of Jacob and Messiah's Light, the son of Karl Hirsch, the son of Abraham, the son of Hillel, when the LORD first drew him out from the nations and inclined his spirit to seek after the LORD. It first came when he was about 30 [1995/6] years of age, saying expressly: \"Thou shalt surely be my witness to Israel.\""},
	}
	bibleData.Verses["en"]["BEN.1"] = ben11EN
}
