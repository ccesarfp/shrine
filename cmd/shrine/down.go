package main

import (
	"github.com/ccesarfp/shrine/pkg/util"
	"github.com/spf13/cobra"
	"log"
	"os"
	"syscall"
)

var force bool

// down - represents the down command
func down() *cobra.Command {
	downCmd := &cobra.Command{
		Use:   "down",
		Short: "Stop server",
		Long:  `Ends running the Shrine server.`,
		Run: func(cmd *cobra.Command, args []string) {
			processList, err := util.FindProcess(ProcessName)
			if err != nil {
				log.Fatal(err)
			}

			signal := os.Interrupt
			if force {
				signal = syscall.SIGTERM
			}

			for _, p := range processList {
				if p != nil {
					_, err = util.SendSignal(p, signal)
					if err != nil {
						log.Fatalln(err)
					}
					log.Println(p.Pid, "off")
				}
			}

		},
	}

	downCmd.Flags().BoolVarP(&force, "force", "f", false, "force stop server")
	return downCmd
}
