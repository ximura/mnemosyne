package main

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ximura/mnemosyne/internal/parser"
	"github.com/ximura/mnemosyne/internal/syncstore"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync ChatGPT conversations to Notion",
	RunE: func(cmd *cobra.Command, args []string) error {
		start := time.Now()

		exportPath := viper.GetString("chatgpt_export_path")
		notionKey := viper.GetString("notion_api_key")
		dbID := viper.GetString("database_id")

		if exportPath == "" || notionKey == "" || dbID == "" {
			return fmt.Errorf("config incomplete, check notion_api_key, database_id, and chatgpt_export_path")
		}

		logger.Info("Starting sync",
			slog.String("export_path", exportPath),
			slog.String("database_id", dbID),
		)

		mng := syncstore.NewFileManager()
		_, err := mng.Load()
		if err != nil {
			return fmt.Errorf("failed to load previous state: %w", err)
		}

		// TODO: implement parsing, tagging, and Notion upload
		chats, err := parser.LoadChats(exportPath)
		if err != nil {
			return fmt.Errorf("failed to parse chats: %w", err)
		}

		logger.Info("Parsed chats", slog.Int("count", len(chats)))

		logger.Info("sync completed",
			slog.Duration("elapsed", time.Since(start)),
			slog.Int("chats", len(chats)),
		)
		return nil
	},
}
