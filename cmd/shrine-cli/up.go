package main

import (
	"github.com/ccesarfp/shrine/internal/config/application"
	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start server",
	Long:  `Start Shrine server.`,
	Run: func(cmd *cobra.Command, args []string) {

		app = application.New()
		app.SetupServer()
		app.Up()

	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
