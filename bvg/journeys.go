package bvg

import (
	"time"
)

type JourneysQueryParams struct {
	Departure        time.Time `url:"departure,omitempty"`
	Arrival          time.Time `url:"arrival,omitempty"`
	EarlierThan      string    `url:"earlierThan,omitempty"`
	LaterThan        string    `url:"laterThan,omitempty"`
	Results          int       `url:"results,omitempty"`
	Stopovers        bool      `url:"stopovers,omitempty"`
	Transfers        int       `url:"transfers,omitempty"`
	TransferTime     int       `url:"transferTime,omitempty"`
	Accessibility    string    `url:"accessibility,omitempty"`
	Bike             bool      `url:"bike,omitempty"`
	StartWithWalking bool      `url:"startWithWalking,omitempty"`
	WalkingSpeed     string    `url:"walkingSpeed,omitempty"`
	Tickets          bool      `url:"tickets,omitempty"`
	Polylines        bool      `url:"polylines,omitempty"`
	SubStops         bool      `url:"subStops,omitempty"`
	Entrances        bool      `url:"entrances,omitempty"`
	Remarks          bool      `url:"remarks,omitempty"`
	ScheduledDays    bool      `url:"scheduledDays,omitempty"`
	Language         string    `url:"language,omitempty"`
	Suburban         bool      `url:"suburban,omitempty"`
	Subway           bool      `url:"subway,omitempty"`
	Tram             bool      `url:"tram,omitempty"`
	Bus              bool      `url:"bus,omitempty"`
	Ferry            bool      `url:"ferry,omitempty"`
	Express          bool      `url:"express,omitempty"`
	Regional     bool      `url:"regional,omitempty"`
	Pretty           bool      `url:"pretty,omitempty"`
}

func (c *Client) Journeys(params *JourneysQueryParams) (*JourneysResponse, error) {
	var res JourneysResponse
	err := c.sendRequest(journeys(), params, &res)
	return &res, err
}

type JourneysRefQueryParams struct {
	Stopovers     bool   `url:"stopovers,omitempty"`
	Tickets       bool   `url:"tickets,omitempty"`
	Polylines     bool   `url:"polylines,omitempty"`
	SubStops      bool   `url:"subStops,omitempty"`
	Entrances     bool   `url:"entrances,omitempty"`
	Remarks       bool   `url:"remarks,omitempty"`
	ScheduledDays bool   `url:"scheduledDays,omitempty"`
	Language      string `url:"language,omitempty"`
	Pretty        bool   `url:"pretty,omitempty"`
}

func (c *Client) JourneysRef(ref string, params *JourneysRefQueryParams) (*Journey, error) {
	var res Journey
	err := c.sendRequest(journeysRef(ref), params, &res)
	return &res, err
}
