package twitch

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
)

// GetClipsInput represents the query string parameters to get clips
// https://dev.twitch.tv/docs/api/reference#get-clips
type GetClipsInput struct {
	// Cursor for forward pagination: tells the server where to start fetching the next set of results, in a multi-page response. This applies only to queries specifying broadcaster_id or game_id. The cursor value specified here is from the pagination response field of a prior query.
	After string `url:"after"`
	// Cursor for backward pagination: tells the server where to start fetching the next set of results, in a multi-page response. This applies only to queries specifying broadcaster_id or game_id. The cursor value specified here is from the pagination response field of a prior query.
	Before string `url:"before"`
	// ID of the broadcaster for whom clips are returned. The number of clips returned is determined by the first query-string parameter (default: 20). Results are ordered by view count.
	BroadcasterID string `url:"broadcaster_id"`
	// Ending date/time for returned clips, in RFC3339 format. (Note that the seconds value is ignored.) If this is specified, started_at also must be specified; otherwise, the time period is ignored.
	EndedAt string `url:"ended_at"`
	// Maximum number of objects to return. Maximum: 100. Default: 20.
	First uint `url:"first"`
	// ID of the game for which clips are returned. The number of clips returned is determined by the first query-string parameter (default: 20). Results are ordered by view count.
	GameID string `url:"game_id"`
	// ID of the clip being queried. Limit: 100.
	ID string `url:"id"`
	// Starting date/time for returned clips, in RFC3339 format. (Note that the seconds value is ignored.) If this is specified, ended_at also should be specified; otherwise, the ended_at date/time will be 1 week after the started_at value.
	StartedAt string `url:"started_id"`
}

// ClipData represents the data in a single clip
type ClipData struct {
	// ID of the clip being queried.
	ID string `json:"id"`
	// User ID of the stream from which the clip was created.
	BroadcasterID string `json:"broadcaster_id"`
	// Display name corresponding to broadcaster_id.
	BroadcasterName string `json:"broadcaster_name"`
	// Date when the clip was created.
	CreatedAt string `json:"created_at"`
	// ID of the user who created the clip.
	CreatorID string `json:"creator_id"`
	// Display name corresponding to creator_id.
	CreatorName string `json:"creator_name"`
	// URL to embed the clip.
	EmbedURL string `json:"embed_url"`
	// ID of the game assigned to the stream when the clip was created.
	GameID string `json:"game_id"`
	// Language of the stream from which the clip was created.
	Language string `json:"language"`
	// URL of the clip thumbnail.
	ThumbnailURL string `json:"thumbnail_url"`
	// Title of the clip.
	Title string `json:"title"`
	// URL where the clip can be viewed.
	URL string `json:"url"`
	// ID of the video from which the clip was created.
	VideoID string `json:"video_id"`
	// Number of times the clip has been viewed.
	ViewCount uint `json:"view_count"`
}

type clips struct {
	Data       []ClipData `json:"data"`
	Pagination `json:"pagination"`
}

// GetClips will get a list of clips
// The url query parameters are defined by the GetClipsInput struct
func (c Client) GetClips(i GetClipsInput) ([]ClipData, *Pagination, error) {
	// since first, when uninitialized is 0, we have to set it to the default value
	if i.First == 0 {
		i.First = 20
	}

	var uri *url.URL

	uri, err := url.Parse(baseURL + getClipsEndpoint)
	if err != nil {
		return nil, nil, err
	}

	uri.RawQuery = parseURLQuery(i).Encode()
	res, err := c.doRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()

	o := clips{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	json.Unmarshal(body, &o)
	return o.Data, &o.Pagination, nil
}
