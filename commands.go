package main

import (
	"sort"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"patmaeder.com/bvg-smartscreen/bvg"
)

const (
	refreshRate = 1 * time.Minute
	stationID   = "900009202"
	stationName = "U Osloer Straße"
)

func (m model) fetchData() tea.Cmd {
	return func() tea.Msg {
		res, err := m.client.StopsIdDepartures(stationID, &bvg.StopsIdDeparturesQueryParams{
			Duration: 60,
			Results:  60,
			Language: "de",
			Suburban: true,
			Subway:   true,
			Tram:     true,
			Bus:      true,
			Remarks:  true,
		})
		if err != nil {
			return errMsg{err}
		}

		var filtered []bvg.Departure
		for _, d := range res.Departures {
			if (d.Line.Name == "128" || d.Line.Name == "125") && strings.Contains(strings.ToLower(d.Direction), "osloer") {
				continue
			}

			t := d.When
			if t.IsZero() {
				t = d.PlannedWhen
			}
			if time.Until(t) > 5*time.Minute {
				filtered = append(filtered, d)
			}
		}

		productOrder := map[string]int{
			"suburban": 1,
			"subway":   2,
			"tram":     3,
			"bus":      4,
		}

		sort.Slice(filtered, func(i, j int) bool {
			pi := productOrder[filtered[i].Line.Product]
			pj := productOrder[filtered[j].Line.Product]
			if pi != pj {
				return pi < pj
			}
			ti := filtered[i].When
			if ti.IsZero() {
				ti = filtered[i].PlannedWhen
			}
			tj := filtered[j].When
			if tj.IsZero() {
				tj = filtered[j].PlannedWhen
			}
			return ti.Before(tj)
		})

		nauenerArrivals := make(map[string]time.Time)
		notifMap := make(map[string]bool)

		for _, dep := range filtered {
			if dep.Line.Name == "U9" {
				trip, err := m.client.TripsId(dep.TripID, &bvg.TripsIdQueryParams{Stopovers: true})
				if err == nil {
					for _, so := range trip.Stopovers {
						if so.Stop.ID == "900009201" {
							arr := so.Arrival
							if arr.IsZero() {
								arr = so.Departure
							}
							nauenerArrivals[dep.TripID] = arr
							break
						}
					}
				}
			}

			for _, r := range dep.Remarks {
				text := r.Summary
				if text == "" {
					text = r.Text
				}
				if (r.Type == "warning" || r.Type == "status") && text != "" {
					if !strings.Contains(strings.ToLower(text), "kurzzug") {
						notifMap[strings.TrimSpace(text)] = true
					}
				}
			}
		}

		var notifications []string
		for k := range notifMap {
			notifications = append(notifications, k)
		}
		sort.Strings(notifications)

		return dataMsg{
			departures:      filtered,
			nauenerArrivals: nauenerArrivals,
			notifications:   notifications,
			lastUpdate:      time.Now(),
		}
	}
}

func (m model) fetchTick() tea.Cmd {
	return tea.Tick(refreshRate, func(t time.Time) tea.Msg {
		return fetchTickMsg(t)
	})
}

func (m model) scrollTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return scrollTickMsg(t)
	})
}
