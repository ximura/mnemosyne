package main

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	logger  *slog.Logger
)

func main() {
	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	rootCmd := &cobra.Command{
		Use:   "mnemo",
		Short: "Mnemosyne — Sync your ChatGPT conversations to Notion",
		Long: `Mnemosyne is a CLI tool that syncs your ChatGPT conversations
into a Notion database, organizing them by tags and topics.`,
	}

	// Config flag
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mnemo.yaml)")
	cobra.OnInitialize(initConfig)

	// Add commands
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(previewCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.Error("command failed", "error", err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			logger.Error("cannot detect home directory", "error", err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".mnemo")
	}

	if err := viper.ReadInConfig(); err == nil {
		logger.Info("config loaded",
			slog.String("filename", viper.ConfigFileUsed()),
		)
	} else {
		logger.Info("⚠️  No config file found — using defaults.")
	}
}
