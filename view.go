package main

import (
	"fmt"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	"patmaeder.com/bvg-smartscreen/bvg"
)

func (m model) View() tea.View {
	if m.err != nil {
		content := errorStyle.Render(fmt.Sprintf("⚠ Fehler: %v", m.err)) + "\n\n[r] Erneut versuchen  [q] Beenden"
		v := tea.NewView(content)
		v.AltScreen = true
		return v
	} else if len(m.departures) == 0 {
		content := "  Lade Abfahrtsdaten...\n"
		v := tea.NewView(content)
		v.AltScreen = true
		return v
	}

	var ubahnDeps, tramDeps, busDeps []bvg.Departure
	for _, dep := range m.departures {
		switch dep.Line.Product {
		case "subway":
			ubahnDeps = append(ubahnDeps, dep)
		case "tram":
			tramDeps = append(tramDeps, dep)
		case "bus":
			busDeps = append(busDeps, dep)
		}
	}

	halfWidth := m.width / 2
	if halfWidth < 20 {
		halfWidth = 20
	}

	availableHeight := m.height - 4
	ubahnHeight := int(float64(availableHeight) * 0.45)
	bottomHeight := availableHeight - ubahnHeight

	now := time.Now()
	days := []string{"So", "Mo", "Di", "Mi", "Do", "Fr", "Sa"}
	day := days[now.Weekday()]

	timeAndDateStr := fmt.Sprintf("%s, %02d.%02d.  |  %s ", day, now.Day(), now.Month(), now.Format("15:04:05"))
	stationStr := " S+U Osloer Straße"

	headerSpace := m.width - lipgloss.Width(stationStr) - lipgloss.Width(timeAndDateStr)
	if headerSpace < 0 {
		headerSpace = 0
	}

	headerRow := headerRowStyle.Width(m.width).Render(
		stationStr +
			strings.Repeat(" ", headerSpace) +
			timeAndDateStr,
	)

	ubahnTile := m.renderTile("U-Bahn", ubahnDeps, tileHeaderUbahn, m.width, ubahnHeight)
	tramTile := m.renderTile("Tram", tramDeps, tileHeaderTram, halfWidth, bottomHeight)
	busTile := m.renderTile("Bus", busDeps, tileHeaderBus, halfWidth, bottomHeight)

	bottomRow := lipgloss.JoinHorizontal(lipgloss.Top, tramTile, busTile)

	var marqueeString string
	if len(m.notifications) > 0 {
		marqueeStr := "   +++   " + strings.Join(m.notifications, "   +++   ")
		runes := []rune(marqueeStr)
		if len(runes) > 0 {
			offset := m.marqueeOffset % len(runes)
			var visibleRunes []rune
			for i := 0; i < m.width; i++ {
				idx := (offset + i) % len(runes)
				visibleRunes = append(visibleRunes, runes[idx])
			}
			marqueeString = marqueeStyle.Render(string(visibleRunes))
		}
	}

	finalLayout := lipgloss.JoinVertical(lipgloss.Left,
		headerRow,
		"",
		ubahnTile,
		bottomRow,
		marqueeString,
	)

	v := tea.NewView(finalLayout)
	v.AltScreen = true
	return v
}

type depGroup struct {
	LineName   string
	Product    string
	Direction  string
	Departures []bvg.Departure
}

func isSimilarDirection(dir1, dir2 string) bool {
	if dir1 == dir2 {
		return true
	}

	normalize := func(s string) string {
		s = strings.ToLower(s)
		s = strings.ReplaceAll(s, "s+u ", "")
		s = strings.ReplaceAll(s, "s ", "")
		s = strings.ReplaceAll(s, "u ", "")
		if idx := strings.Index(s, ","); idx != -1 {
			s = s[:idx]
		}
		return strings.TrimSpace(s)
	}

	return normalize(dir1) == normalize(dir2)
}

func (m model) renderTile(title string, deps []bvg.Departure, headerStyle lipgloss.Style, tileWidth, tileHeight int) string {
	var header strings.Builder
	header.WriteString(headerStyle.Width(tileWidth - 4).Render(title))
	header.WriteString("\n\n")

	if len(deps) == 0 {
		return tileStyle.
			Width(tileWidth).MaxWidth(tileWidth).
			Height(tileHeight).MaxHeight(tileHeight).
			Render(header.String() + lipgloss.NewStyle().Foreground(colorGray).Italic(true).Render("  Keine Abfahrten  "))
	}

	var groups []depGroup
	for _, d := range deps {
		added := false
		for i, g := range groups {
			if g.LineName == d.Line.Name && isSimilarDirection(g.Direction, d.Direction) {
				groups[i].Departures = append(groups[i].Departures, d)
				if len(d.Direction) > len(g.Direction) {
					groups[i].Direction = d.Direction
				}
				added = true
				break
			}
		}
		if !added {
			groups = append(groups, depGroup{
				LineName:   d.Line.Name,
				Product:    d.Line.Product,
				Direction:  d.Direction,
				Departures: []bvg.Departure{d},
			})
		}
	}

	var content strings.Builder
	for _, g := range groups {
		lineBox := getLineStyle(g.Product, g.LineName).Render(g.LineName)
		dirText := lipgloss.NewStyle().Foreground(colorBlack).Bold(true).Render(g.Direction)

		content.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, lineBox, "  ", dirText))
		content.WriteString("\n")

		var timeStrs []string
		for i, dep := range g.Departures {
			if i >= 3 {
				break
			}

			actualTime := dep.When
			if actualTime.IsZero() {
				actualTime = dep.PlannedWhen
			}

			minutesLeft := int(time.Until(actualTime).Minutes())
			if minutesLeft < 0 {
				continue
			}

			delayStr := ""
			if dep.Delay > 0 {
				delayStr = lipgloss.NewStyle().Foreground(colorRed).Render(fmt.Sprintf("+%d", dep.Delay/60))
			}

			baseTimeStr := fmt.Sprintf("%d min", minutesLeft)
			if len(timeStrs) == 0 {
				baseTimeStr = firstTimeStyle.Render(baseTimeStr)
			} else {
				baseTimeStr = subsequentTimeStyle.Render(baseTimeStr)
			}

			if delayStr != "" {
				baseTimeStr += " " + delayStr
			}

			if g.LineName == "U9" && len(timeStrs) == 0 {
				if arr, ok := m.nauenerArrivals[dep.TripID]; ok {
					baseTimeStr += " " + arrivalStyle.Render(fmt.Sprintf("(Nauener: %s)", arr.Format("15:04")))
				}
			}

			for _, r := range dep.Remarks {
				if strings.Contains(strings.ToLower(r.Text), "kurzzug") {
					baseTimeStr += lipgloss.NewStyle().Foreground(colorGray).Render(" ")
					break
				}
			}

			if dep.Direction != g.Direction {
				shortDest := lipgloss.NewStyle().Foreground(colorGray).Render(fmt.Sprintf(" (%s)", dep.Direction))
				baseTimeStr += shortDest
			}

			timeStrs = append(timeStrs, baseTimeStr)
		}

		if len(timeStrs) > 0 {
			timesLine := "    " + strings.Join(timeStrs, lipgloss.NewStyle().Foreground(colorGray).Render("  ·  "))
			content.WriteString(timesLine + "\n\n")
		}
	}

	contentStr := content.String()
	lines := strings.Split(strings.TrimRight(contentStr, "\n"), "\n")
	maxContentHeight := tileHeight - 3
	if maxContentHeight < 1 {
		maxContentHeight = 1
	}

	if len(lines) > maxContentHeight {
		start := m.scrollOffset % ((len(lines) - maxContentHeight) + 1)
		lines = lines[start : start+maxContentHeight]
	}

	result := header.String() + strings.Join(lines, "\n")

	return tileStyle.
		Width(tileWidth).MaxWidth(tileWidth).
		Height(tileHeight).MaxHeight(tileHeight).
		Render(result)
}
