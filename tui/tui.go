package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type ServiceStatus struct {
	Name        string
	Status      string
	Description string
	UpdatedAt   time.Time
}

type Model struct {
	services []ServiceStatus
	selected int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	s := "Cloud Status Dashboard\n"
	s += "───────────────────────────────\n"

	for i, svc := range m.services {
		cursor := " "
		if i == m.selected {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %-10s %s\n", cursor, svc.Name, svc.Status)
	}

	s += "\nDetails:\n───────────────────────────────\n"
	sel := m.services[m.selected]
	s += fmt.Sprintf("Service: %s\nStatus : %s\nUpdated: %s\nInfo   : %s\n",
		sel.Name, sel.Status, sel.UpdatedAt.Format(time.RFC822), sel.Description)

	return s
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.selected > 0 {
				m.selected--
			}
		case "down":
			if m.selected < len(m.services)-1 {
				m.selected++
			}
		}
	}

	return m, nil
}
