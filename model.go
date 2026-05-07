package main

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"patmaeder.com/bvg-smartscreen/bvg"
)


type dataMsg struct {
	departures      []bvg.Departure
	nauenerArrivals map[string]time.Time
	notifications   []string
	lastUpdate      time.Time
}

type fetchTickMsg time.Time
type scrollTickMsg time.Time
type errMsg struct{ err error }

type model struct {
	client          *bvg.Client
	departures      []bvg.Departure
	nauenerArrivals map[string]time.Time
	notifications   []string
	lastUpdated     time.Time
	err             error
	loading         bool
	width           int
	height          int
	scrollOffset    int
	marqueeOffset   int
	tickCount       int
}

func initialModel() model {
	return model{
		client:          bvg.NewClient(),
		nauenerArrivals: make(map[string]time.Time),
		loading:         true,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.fetchData(),
		m.fetchTick(),
		m.scrollTick(),
	)
}
