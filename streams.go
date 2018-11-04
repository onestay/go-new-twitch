package twitch

import (
	"encoding/json"
)

// GetStreamInput represents the available URL query parameters for the stream endpoint as defined by the Twitch API
// https://dev.twitch.tv/docs/api/reference/#get-streams
type GetStreamInput struct {
	After       string
	CommunityID string
	GameID      string
	Language    string
	UserID      string
	UserLogin   string
	Limit       int
}

type urlFields struct {
	After       string `url:"after,omitempty"`
	CommunityID string `url:"community_id,omitempty"`
	First       int    `url:"first,omitempty"`
	GameID      string `url:"game_id,omitempty"`
	Language    string `url:"language,omitempty"`
	UserID      string `url:"user_id,omitempty"`
	UserLogin   string `url:"user_login,omitempty"`
	Limit       int    `url:"-"`
}

// GetStreams will return the streams specified by GetStreamInput
func (c *Client) GetStreams(i GetStreamInput) ([]Stream, error) {
	u := urlFields{
		After:       i.After,
		CommunityID: i.CommunityID,
		GameID:      i.GameID,
		Language:    i.Language,
		UserID:      i.UserID,
		UserLogin:   i.UserLogin,
		Limit:       i.Limit,
	}

	if u.Limit < 100 {
		u.First = u.Limit
		u.Limit = 0
	} else {
		u.First = 100
		u.Limit -= 100
	}

	uri, err := c.buildURL(streamsEndpoint, u)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(uri.String(), "GET", nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var streamsSlice []Stream
	var streams twitchResponseStream

	json.NewDecoder(res.Body).Decode(&streams)

	streamsSlice = append(streamsSlice, streams.Data...)
	for u.Limit > 100 {
		u.First = 100
		u.After = streams.Pagination.Cursor
		u.Limit -= 100

		uri2, err := c.buildURL(streamsEndpoint, u)
		if err != nil {
			return nil, err
		}

		res, err := c.doRequest(uri2.String(), "GET", nil)
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()

		streams = twitchResponseStream{}
		json.NewDecoder(res.Body).Decode(&streams)

		streamsSlice = append(streamsSlice, streams.Data...)
	}

	if u.Limit != 0 {
		u.First = u.Limit
		u.After = streams.Pagination.Cursor

		finalURI, err := c.buildURL(streamsEndpoint, u)
		if err != nil {
			return nil, err
		}

		res, err = c.doRequest(finalURI.String(), "GET", nil)
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()

		streams = twitchResponseStream{}
		json.NewDecoder(res.Body).Decode(&streams)

		streamsSlice = append(streamsSlice, streams.Data...)
	}

	return streamsSlice, nil
}
