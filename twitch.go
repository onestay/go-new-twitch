package twitch

import (
	"net/http"
	"time"
)

// Client represents the base client to interact with the twitch api
type Client struct {
	httpClient http.Client
	clientID   string
}

// NewClient initiates a new client to use with the twitch api
func NewClient(id string) *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: time.Second * 20,
		},
		clientID: id,
	}
}
