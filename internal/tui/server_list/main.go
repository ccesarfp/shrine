package server_list

import (
	"github.com/ccesarfp/shrine/pkg/util"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
	"os"
	"strconv"
	"syscall"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)
	force    bool
)

type item struct {
	title, desc string
}

func NewItem(title string, desc string) item {
	return item{
		title: title,
		desc:  desc,
	}
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func NewModel(list list.Model, title string, forceFlag bool) model {
	force = forceFlag
	m := model{
		list: list,
	}
	m.list.Title = title
	return m
}

type model struct {
	list list.Model
}

func (m model) List() list.Model {
	return m.list
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {

			// Getting selected item
			itemSelected := m.list.SelectedItem()

			if itemSelected != nil {
				// Converting PID string to int
				pid, err := strconv.Atoi(itemSelected.FilterValue())
				if err != nil {
					log.Fatalln(err)
				}

				signal := os.Interrupt
				if force {
					signal = syscall.SIGTERM
				}

				// Finding process
				process, err := os.FindProcess(pid)
				if err != nil {
					log.Fatalln(err)
				}
				// Sending signal
				_, err = util.SendSignal(process, signal)
				if err != nil {
					log.Fatalln(err)
				}
			}

			tea.ClearScreen()

			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}
