package main

import (
	"fmt"
	"github.com/ccesarfp/shrine/internal/config/application"
	"github.com/ccesarfp/shrine/pkg/util"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var app *application.Application

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start server",
	Long:  `Start the Shrine server.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Checking the number of application processes, if it more than 0, do not start the server
		c, _ := util.CountProcess(ProcessName)
		if c <= 1 {
			app = application.New()
			app.SetupServer()

			// Setting events
			go down()
			go downBrutally()

			app.Up()
		}
		if c > 1 {
			fmt.Println("A process is already running")
		}

	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}

// Event to stop server
func down() {
	var downChan = make(chan os.Signal, 1)
	defer close(downChan)
	signal.Notify(downChan, os.Interrupt)
	<-downChan
	app.Down()
}

// Event to force stop server
func downBrutally() {
	var downForceChan = make(chan os.Signal, 1)
	defer close(downForceChan)
	signal.Notify(downForceChan, syscall.SIGTERM)
	<-downForceChan
	app.DownBrutally()
}
