package main

import (
	"context"
	"github.com/spf13/cobra"
)

var ProcessName string = "shrine"

func Execute() error {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:     "shrine",
		Version: "1.2",
		Short:   "Authentication microservice that enables the rapid and secure creation of JWT tokens.",
		Long:    `Authentication microservice that enables the rapid and secure creation of JWT tokens.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	rootCmd.AddCommand(createKey())
	rootCmd.AddCommand(up())
	rootCmd.AddCommand(status())
	rootCmd.AddCommand(down())
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
	return rootCmd.ExecuteContext(context.Background())
}
