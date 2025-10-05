package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Chat struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	CreatedAt string   `json:"created_at"`
	Tags      []string `json:"-"`
}

// LoadChats reads all JSON files in a folder and parses them
func LoadChats(folder string) ([]Chat, error) {
	var chats []Chat
	err := filepath.WalkDir(folder, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".json" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read file %s: %w", path, err)
		}

		var c struct {
			Conversations []Chat `json:"conversations"`
		}
		if err := json.Unmarshal(data, &c); err != nil {
			return fmt.Errorf("parse %s: %w", path, err)
		}

		chats = append(chats, c.Conversations...)
		return nil
	})

	return chats, err
}
