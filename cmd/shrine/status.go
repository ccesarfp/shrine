package main

import (
	"fmt"
	"github.com/ccesarfp/shrine/internal/config/application"
	"github.com/ccesarfp/shrine/internal/enum/status"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status server",
	Long:  `Show status server.`,
	Run: func(cmd *cobra.Command, args []string) {
		servers, err := application.Read()
		if err != nil {
			log.Fatalln(err)
		}

		var s status.Status
		fmt.Println("PID", " Name", "  Address", "     Runtime", "      Status")
		for _, server := range servers {
			fmt.Println(server.Pid, server.Name, server.Address, time.Since(server.StartTime), s.String(server.Status))
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
