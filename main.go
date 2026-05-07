package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "r" {
			m.loading = true
			return m, m.fetchData()
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case dataMsg:
		m.departures = msg.departures
		m.nauenerArrivals = msg.nauenerArrivals
		m.notifications = msg.notifications
		m.lastUpdated = msg.lastUpdate
		m.loading = false
		m.err = nil

	case fetchTickMsg:
		return m, tea.Batch(m.fetchData(), m.fetchTick())

	case scrollTickMsg:
		m.tickCount++
		m.marqueeOffset++
		if m.tickCount%13 == 0 {
			m.scrollOffset++
		}
		return m, m.scrollTick()

	case errMsg:
		m.err = msg.err
		m.loading = false
	}

	return m, nil
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
