package main

import (
	"github.com/ccesarfp/shrine/pkg/util"
	"github.com/spf13/cobra"
	"log"
	"os"
	"syscall"
)

var force bool

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop server",
	Long:  `Ends running the Shrine server.`,
	Run: func(cmd *cobra.Command, args []string) {
		p, err := util.FindProcess(ProcessName)
		if err != nil {
			log.Fatal(err)
		}

		signal := os.Interrupt
		if force {
			signal = syscall.SIGTERM
		}

		_, err = util.SendSignal(p, signal)
		if err != nil {
			log.Fatalln(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(downCmd)
	downCmd.Flags().BoolVarP(&force, "force", "f", false, "force stop server")
}
