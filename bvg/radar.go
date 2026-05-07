package bvg

type RadarQueryParams struct {
	North     float64 `url:"north,required"`
	West      float64 `url:"west,required"`
	South     float64 `url:"south,required"`
	East      float64 `url:"east,required"`
	Results   int     `url:"results,omitempty"`
	Duration  int     `url:"duration,omitempty"`
	Frames    int     `url:"frames,omitempty"`
	Polylines bool    `url:"polylines,omitempty"`
	Language  string  `url:"language,omitempty"`
	Pretty    bool    `url:"pretty,omitempty"`
}

func (c *Client) Radar(params *RadarQueryParams) (*RadarResponse, error) {
	var res RadarResponse
	err := c.sendRequest(radar(), params, &res)
	return &res, err
}
