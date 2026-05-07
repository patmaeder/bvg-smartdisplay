package bvg

import (
	"time"
)

type StopsQueryParams struct {
	Query      string `url:"query,omitempty"`
	Results    int    `url:"results,omitempty"`
	Fuzzy      bool   `url:"fuzzy,omitempty"`
	Completion bool   `url:"completion,omitempty"`
}

func (c *Client) Stops(params *StopsQueryParams) (*[]Location, error) {
	var res []Location
	err := c.sendRequest(stops(), params, &res)
	return &res, err
}

type StopsReachableFromQueryParams struct {
	Latitude     float64 `url:"latitude,required"`
	Longitude    float64 `url:"longitude,required"`
	Address      string  `url:"address,required"`
	When         string  `url:"when,omitempty"`
	MaxTransfers int     `url:"maxTransfers,omitempty"`
	MaxDuration  int     `url:"maxDuration,omitempty"`
	Language     string  `url:"language,omitempty"`
	Suburban     bool    `url:"suburban,omitempty"`
	Subway       bool    `url:"subway,omitempty"`
	Tram         bool    `url:"tram,omitempty"`
	Bus          bool    `url:"bus,omitempty"`
	Ferry        bool    `url:"ferry,omitempty"`
	Express      bool    `url:"express,omitempty"`
	Regional     bool    `url:"regional,omitempty"`
	Pretty       bool    `url:"pretty,omitempty"`
}

func (c *Client) StopsReachableFrom(params *StopsReachableFromQueryParams) (*ReachableResponse, error) {
	var res ReachableResponse
	err := c.sendRequest(stopsReachableFrom(), params, &res)
	return &res, err
}

type StopsIdQueryParams struct {
	LinesOfStops bool   `url:"linesOfStops,omitempty"`
	Language     string `url:"language,omitempty"`
	Pretty       bool   `url:"pretty,omitempty"`
}

func (c *Client) StopsId(id string, params *StopsIdQueryParams) (*Location, error) {
	var res Location
	err := c.sendRequest(stopsId(id), params, &res)
	return &res, err
}

type StopsIdDeparturesQueryParams struct {
	When         time.Time `url:"when,omitempty"`
	Direction    string    `url:"direction,omitempty"`
	Duration     int       `url:"duration,omitempty"`
	Results      int       `url:"results,omitempty"`
	LinesOfStops bool      `url:"linesOfStops,omitempty"`
	Remarks      bool      `url:"remarks,omitempty"`
	Language     string    `url:"language,omitempty"`
	Suburban     bool      `url:"suburban,omitempty"`
	Subway       bool      `url:"subway,omitempty"`
	Tram         bool      `url:"tram,omitempty"`
	Bus          bool      `url:"bus,omitempty"`
	Ferry        bool      `url:"ferry,omitempty"`
	Express      bool      `url:"express,omitempty"`
	Regional     bool      `url:"regional,omitempty"`
	Pretty       bool      `url:"pretty,omitempty"`
}

func (c *Client) StopsIdDepartures(id string, params *StopsIdDeparturesQueryParams) (*DeparturesResponse, error) {
	var res DeparturesResponse
	err := c.sendRequest(stopsIdDepartures(id), params, &res)
	return &res, err
}

type StopsIdArrivalsQueryParams struct {
	When         time.Time `url:"when,omitempty"`
	Direction    string    `url:"direction,omitempty"`
	Duration     int       `url:"duration,omitempty"`
	Results      int       `url:"results,omitempty"`
	LinesOfStops bool      `url:"linesOfStops,omitempty"`
	Remarks      bool      `url:"remarks,omitempty"`
	Language     string    `url:"language,omitempty"`
	Suburban     bool      `url:"suburban,omitempty"`
	Subway       bool      `url:"subway,omitempty"`
	Tram         bool      `url:"tram,omitempty"`
	Bus          bool      `url:"bus,omitempty"`
	Ferry        bool      `url:"ferry,omitempty"`
	Express      bool      `url:"express,omitempty"`
	Regional     bool      `url:"regional,omitempty"`
	Pretty       bool      `url:"pretty,omitempty"`
}

func (c *Client) StopsIdArrivals(id string, params *StopsIdArrivalsQueryParams) (*ArrivalsResponse, error) {
	var res ArrivalsResponse
	err := c.sendRequest(stopsIdArrivals(id), params, &res)
	return &res, err
}
