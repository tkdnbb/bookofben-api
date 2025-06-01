package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var data = make([][]string, 73)

func init() {
	// 加载已有的章节
	chapters := []int{1, 2} // 在这里添加你想要加载的章节号

	for _, chapterNum := range chapters {
		loadChapter(chapterNum)
	}
}

// loadChapter 加载指定章节
func loadChapter(chapterNum int) {
	chapterPath := filepath.Join("internal", "data", "chapters", fmt.Sprintf("chapter%d.json", chapterNum))
	content, err := os.ReadFile(chapterPath)
	if err != nil {
		log.Fatalf("Failed to read chapter%d.json: %v", chapterNum, err)
	}

	var verses []string
	if err := json.Unmarshal(content, &verses); err != nil {
		log.Fatalf("Failed to unmarshal chapter%d.json: %v", chapterNum, err)
	}

	data[chapterNum-1] = verses
	log.Printf("Loaded %d verses for chapter %d", len(verses), chapterNum)
}

// GetChapterVerses 返回指定章节的经文
func GetChapterVerses(chapter int) []string {
	if chapter < 1 || chapter > 73 {
		return nil
	}
	return data[chapter-1]
}

// GetAllData 返回所有数据（用于调试）
func GetAllData() [][]string {
	return data
}
