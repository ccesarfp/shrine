package main

import (
	"fmt"
	"github.com/ccesarfp/shrine/internal/config/application"
	statusEnum "github.com/ccesarfp/shrine/internal/enum/status"
	"github.com/ccesarfp/shrine/internal/tui/status_table"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"time"
)

// status - represents the status command
func status() *cobra.Command {
	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Show status server",
		Long:  `Show status server.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Getting Data
			servers, err := application.Read()
			if err != nil {
				log.Fatalln(err)
			}

			// Have servers running
			if len(servers) > 0 {
				// Creating Columns
				columns := []table.Column{
					{Title: "PID", Width: 8},
					{Title: "Name", Width: 10},
					{Title: "Version", Width: 8},
					{Title: "Address", Width: 14},
					{Title: "Runtime", Width: 18},
					{Title: "Status", Width: 8},
				}

				// Creating Rows
				var rows []table.Row
				for _, server := range servers {
					rows = append(rows, []table.Row{
						{
							strconv.Itoa(server.Pid),
							server.Name,
							server.Version,
							server.Address,
							time.Since(server.StartTime).String(),
							statusEnum.String(server.Status),
						},
					}...)
				}

				// Creating Table
				t := table.New(
					table.WithColumns(columns),
					table.WithRows(rows),
					table.WithFocused(true),
					table.WithHeight(len(rows)),
				)

				// Customizing Table Style
				s := table.DefaultStyles()
				s.Header = s.Header.
					BorderStyle(lipgloss.NormalBorder()).
					BorderForeground(lipgloss.Color("240")).
					BorderBottom(true).
					Bold(false)
				s.Selected = s.Selected.
					Foreground(lipgloss.Color("229")).
					Background(lipgloss.Color("#800F9B")).
					Bold(false)
				t.SetStyles(s)

				// Starting Table
				m := status_table.Model{Table: t}
				if _, err := tea.NewProgram(m).Run(); err != nil {
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
	return statusCmd
}
