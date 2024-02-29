package main

import (
	"github.com/ccesarfp/shrine/internal/config/application"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var app *application.Application

// up - represents the up command
func up() *cobra.Command {
	upCmd := &cobra.Command{
		Use:   "up",
		Short: "Start server",
		Long:  `Start the Shrine server.`,
		Run: func(cmd *cobra.Command, args []string) {

			app = application.New()
			app.S.SetupServer()

			// Setting events
			go downServer()
			go downServerBrutally()

			app.Up()

		},
	}

	return upCmd
}

// Event to stop server
func downServer() {
	var downChan = make(chan os.Signal, 1)
	defer close(downChan)
	signal.Notify(downChan, os.Interrupt)
	<-downChan
	app.Down()
}

// Event to force stop server
func downServerBrutally() {
	var downForceChan = make(chan os.Signal, 1)
	defer close(downForceChan)
	signal.Notify(downForceChan, syscall.SIGTERM)
	<-downForceChan
	app.DownBrutally()
}
