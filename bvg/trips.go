package bvg

import (
	"time"
)

type TripsQueryParams struct {
	Query                string    `url:"query,required"`
	When                 time.Time `url:"when,omitempty"`
	FromWhen             time.Time `url:"fromWhen,omitempty"`
	UntilWhen            time.Time `url:"untilWhen,omitempty"`
	OnlyCurrentlyRunning bool      `url:"onlyCurrentlyRunning,omitempty"`
	CurrentlyStoppingAt  string    `url:"currentlyStoppingAt,omitempty"`
	LineName             string    `url:"lineName,omitempty"`
	OperatorNames        string    `url:"operatorNames,omitempty"`
	Stopovers            bool      `url:"stopovers,omitempty"`
	Remarks              bool      `url:"remarks,omitempty"`
	SubStops             bool      `url:"subStops,omitempty"`
	Entrances            bool      `url:"entrances,omitempty"`
	Language             string    `url:"language,omitempty"`
	Suburban             bool      `url:"suburban,omitempty"`
	Subway               bool      `url:"subway,omitempty"`
	Tram                 bool      `url:"tram,omitempty"`
	Bus                  bool      `url:"bus,omitempty"`
	Ferry                bool      `url:"ferry,omitempty"`
	Express              bool      `url:"express,omitempty"`
	Regional             bool      `url:"regional,omitempty"`
	Pretty               bool      `url:"pretty,omitempty"`
}

func (c *Client) Trips(params *TripsQueryParams) (*TripsResponse, error) {
	var res TripsResponse
	err := c.sendRequest(trips(), params, &res)
	return &res, err
}

type TripsIdQueryParams struct {
	Stopovers bool   `url:"stopovers,omitempty"`
	Remarks   bool   `url:"remarks,omitempty"`
	Polyline  bool   `url:"polyline,omitempty"`
	Language  string `url:"language,omitempty"`
	Pretty    bool   `url:"pretty,omitempty"`
}

func (c *Client) TripsId(id string, params *TripsIdQueryParams) (*Trip, error) {
	var res struct {
		Trip                  Trip  `json:"trip"`
		RealtimeDataUpdatedAt int64 `json:"realtimeDataUpdatedAt,omitempty"`
	}
	err := c.sendRequest(tripsId(id), params, &res)
	return &res.Trip, err
}
