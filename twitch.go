package twitch

import (
	"net/http"
)

// Client represents a client to interact with the twitch API
type Client struct {
	ClientID   string
	httpClient *http.Client
}

// NewClient will initialize a new client for the twitch api
func NewClient(cid string) *Client {
	return &Client{
		ClientID:   cid,
		httpClient: &http.Client{},
	}
}
