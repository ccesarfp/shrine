package main

import (
	"github.com/spf13/cobra"
	"os"
)

var ProcessName string = "shrine"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shrine",
	Short: "Authentication microservice that enables the rapid and secure creation of JWT tokens.",
	Long:  `Authentication microservice that enables the rapid and secure creation of JWT tokens.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
