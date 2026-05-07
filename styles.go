package main

import lipgloss "charm.land/lipgloss/v2"

var (
	colorYellow = lipgloss.Color("#F0D722")
	colorWhite  = lipgloss.Color("#FFFFFF")
	colorBlack  = lipgloss.Color("#000000")
	colorGray   = lipgloss.Color("#555555")
	colorRed    = lipgloss.Color("#FF3333")

	colorText = lipgloss.Color("#000000")

	colorUbahn = lipgloss.Color("#0066AD")
	colorTram  = lipgloss.Color("#CC0000")
	colorBus   = lipgloss.Color("#991C61")
	colorU8    = lipgloss.Color("#224F86")
	colorU9    = lipgloss.Color("#F3791D")
	colorBusFg = lipgloss.Color("#511475")

	headerRowStyle = lipgloss.NewStyle().
			Background(colorBlack).
			Foreground(colorWhite).
			Bold(true)

	stationNameStyle = lipgloss.NewStyle().
				Bold(true).
				PaddingLeft(1)

	clockStyle = lipgloss.NewStyle().
			Bold(true).
			PaddingRight(1)

	tileStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#CCCCCC")).
			Padding(0, 1)

	tileHeaderUbahn = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorWhite).
			Background(colorUbahn).
			Padding(0, 1).
			Align(lipgloss.Center)

	tileHeaderTram = tileHeaderUbahn.Background(colorTram)
	tileHeaderBus  = tileHeaderUbahn.Background(colorBus)

	baseLineStyle = lipgloss.NewStyle().
			Bold(true).
			Width(6).
			Align(lipgloss.Center)

	firstTimeStyle = lipgloss.NewStyle().
			Foreground(colorBlack).
			Background(colorYellow).
			Padding(0, 1).
			Bold(true)

	subsequentTimeStyle = lipgloss.NewStyle().
				Foreground(colorGray)

	destStyle = lipgloss.NewStyle().
			Foreground(colorText).
			Width(20)

	minStyle = lipgloss.NewStyle().
			Foreground(colorText).
			Width(8).
			Align(lipgloss.Right)

	errorStyle = lipgloss.NewStyle().
			Foreground(colorRed).
			Bold(true).
			Padding(1)

	arrivalStyle = lipgloss.NewStyle().
			Foreground(colorGray).
			MarginLeft(1)

	marqueeStyle = lipgloss.NewStyle().
			Foreground(colorBlack).
			Background(colorYellow).
			Padding(0, 1).
			Bold(true).
			MarginTop(1)

	updateFooterStyle = lipgloss.NewStyle().
				Foreground(colorGray).
				Italic(true).
				MarginTop(1)
)

func getLineStyle(product string, lineName string) lipgloss.Style {
	switch product {
	case "subway":
		switch lineName {
		case "U8":
			return baseLineStyle.Foreground(colorWhite).Background(colorU8)
		case "U9":
			return baseLineStyle.Foreground(colorWhite).Background(colorU9)
		default:
			return baseLineStyle.Foreground(colorWhite).Background(colorUbahn)
		}
	case "tram":
		return baseLineStyle.Foreground(colorTram).Background(lipgloss.Color("#EEEEEE"))
	case "bus":
		return baseLineStyle.Foreground(colorBusFg).Background(lipgloss.Color("#EEEEEE"))
	default:
		return baseLineStyle.Foreground(colorWhite).Background(colorUbahn)
	}
}
