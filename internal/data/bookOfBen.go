package data

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

var data = make([][]string, 73)

func init() {
	// 加载第1章
	chapter1Path := filepath.Join("internal", "data", "chapters", "chapter1.json")
	content, err := os.ReadFile(chapter1Path)
	if err != nil {
		log.Fatalf("Failed to read chapter1.json: %v", err)
	}

	var verses []string
	if err := json.Unmarshal(content, &verses); err != nil {
		log.Fatalf("Failed to unmarshal chapter1.json: %v", err)
	}

	data[0] = verses
	log.Printf("Loaded %d verses for chapter 1", len(data[0]))
}

// GetChapterVerses 返回指定章节的经文
func GetChapterVerses(chapter int) []string {
	if chapter < 1 || chapter > 73 {
		return nil
	}
	return data[chapter-1] // chapter是1-based，数组是0-based
}

// GetAllData 返回所有数据（用于调试）
func GetAllData() [][]string {
	return data
}
