package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview parsed conversations and detected tags",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Previewing ChatGPT export contents...")
		// TODO: Implement parser preview
		return nil
	},
}
