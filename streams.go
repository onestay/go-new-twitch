package twitch

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"time"
)

// GetStreamsInput represents the query string parameters to get streams
// https://dev.twitch.tv/docs/api/reference#get-streams
type GetStreamsInput struct {
	// Cursor for forward pagination: tells the server where to start fetching the next set of results, in a multi-page response.
	After string `url:"after"`
	// Cursor for backward pagination: tells the server where to start fetching the next set of results, in a multi-page response.
	Before string `url:"before"`
	// Returns streams in a specified community ID. You can specify up to 100 IDs.
	CommunityID []string `url:"community_id"`
	// Maximum number of objects to return. Maximum: 100. Default: 20.
	First int `url:"first"`
	// Returns streams broadcasting a specified game ID. You can specify up to 100 IDs.
	GameID []string `url:"game_id"`
	// Stream language. You can specify up to 100 languages.
	Language []string `url:"language"`
	// Stream type: "all", "live", "vodcast". Default: "all".
	Type string `url:"type"`
	// Returns streams broadcast by one or more specified user IDs. You can specify up to 100 IDs.
	UserID []string `url:"user_id"`
	// Returns streams broadcast by one or more specified user login names. You can specify up to 100 names.
	UserLogin []string `url:"user_login"`
}

// StreamData represents the data a single stream
type StreamData struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	GameID       string    `json:"game_id"`
	CommunityIds []string  `json:"community_ids"`
	Type         string    `json:"type"`
	Title        string    `json:"title"`
	ViewerCount  int       `json:"viewer_count"`
	StartedAt    time.Time `json:"started_at"`
	Language     string    `json:"language"`
	ThumbnailURL string    `json:"thumbnail_url"`
}

type streams struct {
	Data []StreamData `json:"data"`
}

// GetStreams will get a list of live streams
// The url query parameter are defined by the GetStreamsInput struct
func (c Client) GetStreams(i GetStreamsInput) ([]StreamData, error) {
	// since first, when uninitialized is 0, we have to set it to the default value
	if i.First == 0 {
		i.First = 20
	}

	var uri *url.URL

	uri, err := url.Parse(baseURL + getStreamsEndpoint)
	if err != nil {
		return nil, err
	}

	uri.RawQuery = parseURLQuery(i).Encode()
	res, err := c.doRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	s := streams{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &s)
	return s.Data, nil
}
