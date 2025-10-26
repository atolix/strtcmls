package main

import (
	"fmt"
	"os"
	"time"

	"github.com/atolix/strtcmls/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := tui.NewModel([]tui.ServiceStatus{
		{"AWS", "✅ Operational", "All systems nominal.", time.Now()},
		{"GCP", "🟡 Partial Outage", "Issue in europe-west2.", time.Now()},
		{"Azure", "✅ Operational", "No incidents.", time.Now()},
		{"Cloudflare", "🔴 Major Outage", "Global routing degradation.", time.Now()},
	})

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Errpr running program;", err)
		os.Exit(1)
	}
}
