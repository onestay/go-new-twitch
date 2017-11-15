package twitch

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
)

type users struct {
	Data []UserData `json:"data"`
}

// UserData struct represents a user as defined by the twitch api
type UserData struct {
	ID              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
	OfflineImageURL string `json:"offline_image_url"`
	ViewCount       int    `json:"view_count"`
	Email           string `json:"email"`
}

// GetUsersByID gets information about one or more specified Twitch users. Users are identified by ID.
// Limit are 100 users per request
func (c Client) GetUsersByID(ids ...string) ([]UserData, error) {
	return c.getUsers(ids, "id")
}

// GetUsersByLogin gets information about one or more specified Twitch users. Users are identified by login name.
// Limit are 100 users per request.
func (c Client) GetUsersByLogin(names ...string) ([]UserData, error) {
	return c.getUsers(names, "login")
}

func (c Client) getUsers(names []string, meth string) ([]UserData, error) {
	if len(names) > 100 {
		return nil, errors.New("Can only get up to 100 users per request")
	}

	var uri *url.URL

	uri, err := url.Parse(baseURL + getUsersEndpoint)
	if err != nil {
		return nil, err
	}

	p := url.Values{}
	for _, name := range names {
		p.Add(meth, name)
	}

	uri.RawQuery = p.Encode()

	res, err := c.doRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	u := users{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &u)

	return u.Data, nil
}
