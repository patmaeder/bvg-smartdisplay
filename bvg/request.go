package bvg

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
)

func (c *Client) sendRequest(ep endpoint, params interface{}, result interface{}) error {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return err
	}

	u.Path = u.JoinPath(string(ep)).Path

	queryParams, err := query.Values(params)
	if err != nil {
		return err
	}

	if queryParams.Encode() != "" {
		u.RawQuery = queryParams.Encode()
	}

	resp, err := c.HTTPClient.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(result)
}
