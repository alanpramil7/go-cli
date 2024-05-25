package utils

import (
	"strconv"

	"github.com/alanpramil7/go-time-tracker/db"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
			// case "enter":
			// 	return m, tea.Batch(
			// 		tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			// 	)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

// Make rasks of []db.Tasks type to []table.Row
func formatedTasksToRows(tasks []db.Tasks) []table.Row {
	rows := make([]table.Row, len(tasks))
	for i, task := range tasks {
		rows[i] = table.Row{strconv.Itoa(int(task.ID)), task.Date, task.Time, task.Task}
	}

	return rows
}

// Function to create table
func CreateTable(tasks []db.Tasks) model {
	columns := []table.Column{
		{Title: "ID", Width: 5},
		{Title: "Date", Width: 12},
		{Title: "Time", Width: 8},
		{Title: "Task", Width: 12},
	}

	rows := formatedTasksToRows(tasks)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(5),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return model{t}
}
