package twitch

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/google/go-querystring/query"
)

// GetUsersInput serves as the struct to the GetUsers function
type GetUsersInput struct {
	Login []string `url:"login,omitempty"`
	ID    []string `url:"id,omitempty"`
}

// GetUserByLogin will return a single user specified by the login parameter
func (c *Client) GetUserByLogin(login string) (*User, error) {
	u, err := c.getUsers(GetUsersInput{
		Login: []string{login},
	})

	if err != nil {
		return nil, err
	}

	return &u[0], nil
}

// GetUserByID will return a single user specified by the id parameter
func (c *Client) GetUserByID(id string) (*User, error) {
	u, err := c.getUsers(GetUsersInput{
		ID: []string{id},
	})

	if err != nil {
		return nil, err
	}

	return &u[0], nil

}

// GetUsers will return multiple users specified through the GetUsersInput type
func (c *Client) GetUsers(g GetUsersInput) ([]User, error) {
	if len(g.ID) > 100 || len(g.Login) > 100 {
		return nil, errors.New("max 100 IDs or login names at once")
	}

	return c.getUsers(g)
}

func (c *Client) getUsers(u GetUsersInput) ([]User, error) {
	uri, err := url.Parse(baseURL + usersEndpoint)
	if err != nil {
		return nil, err
	}

	v, err := query.Values(u)
	if err != nil {
		return nil, err
	}

	uri.RawQuery = v.Encode()

	res, err := c.doRequest(uri.String(), "GET", nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var users twitchResponseUser

	json.NewDecoder(res.Body).Decode(&users)

	return users.Data, nil
}
