package main

import (
	"fmt"
	"github.com/ccesarfp/shrine/internal/config/application"
	statusEnum "github.com/ccesarfp/shrine/internal/enum/status"
	"github.com/ccesarfp/shrine/internal/tui/server_list"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"time"
)

var force bool

const (
	title           = "Running Servers"
	noServerRunning = "No servers running"
)

// down - represents the down command
func down() *cobra.Command {
	downCmd := &cobra.Command{
		Use:   "down",
		Short: "Stop server",
		Long:  `Ends running the Shrine server.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Getting Servers and Creating Items
			servers, err := application.Read()
			if err != nil {
				log.Fatalln(err)
			}

			// Have servers running
			if len(servers) > 0 {
				var items []list.Item
				for _, server := range servers {
					item := server_list.NewItem(
						strconv.Itoa(server.Pid),
						server.Name+"(v"+server.Version+")"+" - "+server.Address+" - "+time.Since(server.StartTime).String()+" - "+statusEnum.String(server.Status),
					)
					items = append(items, item)
				}

				// Creating List
				m := server_list.NewModel(
					list.New(items, list.NewDefaultDelegate(), 0, 0),
					title,
					force)

				// Starting List
				p := tea.NewProgram(m)
				if _, err := p.Run(); err != nil {
					fmt.Println("Error running program:", err)
					os.Exit(1)
				}
			}

			// Dont have servers running
			if len(servers) <= 0 {
				fmt.Println(noServerRunning)
			}
		},
	}

	downCmd.Flags().BoolVarP(&force, "force", "f", false, "force stop server")
	return downCmd
}
