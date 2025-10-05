package main

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			logger.Error("missing required config fields")
			return fmt.Errorf("config incomplete, check notion_api_key, database_id, and chatgpt_export_path")
		}

		logger.Info("Starting sync",
			slog.String("export_path", exportPath),
			slog.String("database_id", dbID),
		)

		// TODO: implement parsing, tagging, and Notion upload
		logger.Info("sync completed",
			slog.Duration("elapsed", time.Since(start)),
		)
		return nil
	},
}
