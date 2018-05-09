package twitch

import (
	"encoding/json"
	"errors"
	"net/url"
)

// GameData represents the data of a single game
type GameData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BoxArtURL string `json:"box_art_url"`
}

type games struct {
	Data []GameData `json:"data"`
}

// GetGamesByID will get info on games by id.
// At most 100 id values can be specified
func (c Client) GetGamesByID(ids ...string) ([]GameData, error) {
	return c.getGames(ids, "id")
}

// GetGameByName will get info on games by name.
// Game name. The name must be an exact match. For instance, “Pokemon” will not return a list of Pokemon games; instead, query the specific Pokemon game(s) in which you are interested.
// This is a limitation by the twitch api
// At most 100 name values can be specified.
func (c Client) GetGameByName(names ...string) ([]GameData, error) {
	return c.getGames(names, "name")
}

func (c Client) getGames(names []string, meth string) ([]GameData, error) {
	if len(names) > 100 {
		return nil, errors.New("Can only get up to 100 games per request")
	}

	var uri *url.URL
	uri, err := url.Parse(baseURL + getGamesEndpoint)
	if err != nil {
		return nil, err
	}

	p := url.Values{}
	for _, id := range names {
		p.Add(meth, id)
	}

	uri.RawQuery = p.Encode()

	res, err := c.doRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	u := games{}

	err = json.NewDecoder(res.Body).Decode(&u)
	if err != nil {
		return nil, err
	}

	return u.Data, nil
}
