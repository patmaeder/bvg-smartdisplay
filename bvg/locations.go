package bvg

type LocationsQueryParams struct {
	Query        string `url:"query,required"`
	Fuzzy        bool   `url:"fuzzy,omitempty"`
	Results      int    `url:"results,omitempty"`
	Stops        bool   `url:"stops,omitempty"`
	Addresses    bool   `url:"addresses,omitempty"`
	Poi          bool   `url:"poi,omitempty"`
	LinesOfStops bool   `url:"linesOfStops,omitempty"`
	Language     string `url:"language,omitempty"`
	Pretty       bool   `url:"pretty,omitempty"`
}

func (c *Client) Locations(params *LocationsQueryParams) (*[]Location, error) {
	var res []Location
	err := c.sendRequest(locations(), params, &res)
	return &res, err
}

type LocationsNearbyQueryParams struct {
	Latitude     float64 `url:"latitude,required"`
	Longitude    float64 `url:"longitude,required"`
	Results      int     `url:"results,omitempty"`
	Distance     int     `url:"distance,omitempty"`
	Stops        bool    `url:"stops,omitempty"`
	Poi          bool    `url:"poi,omitempty"`
	LinesOfStops bool    `url:"linesOfStops,omitempty"`
	Language     string  `url:"language,omitempty"`
	Pretty       bool    `url:"pretty,omitempty"`
}

func (c *Client) LocationsNearby(params *LocationsNearbyQueryParams) (*[]Location, error) {
	var res []Location
	err := c.sendRequest(locationsNearby(), params, &res)
	return &res, err
}
