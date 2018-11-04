package twitch

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

func (c *Client) doRequest(uri, meth string, body io.ReadCloser) (*http.Response, error) {
	req, err := http.NewRequest(meth, uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Client-ID", c.clientID)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 || res.StatusCode < 200 {
		return nil, c.createTwitchAPIError(res.Body)
	}

	return res, nil
}

func (c *Client) createTwitchAPIError(body io.ReadCloser) error {
	var te twitchError

	json.NewDecoder(body).Decode(&te)

	return &te
}

func (c *Client) buildURL(endpoint string, q interface{}) (*url.URL, error) {
	uri, err := url.Parse(baseURL + endpoint)
	if err != nil {
		return nil, err
	}

	v, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	uri.RawQuery = v.Encode()

	return uri, nil
}
