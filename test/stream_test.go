package test

import (
	"testing"

	"github.com/onestay/go-new-twitch"
)

func TestStreams(t *testing.T) {
	client := twitch.NewClient("")

	s, err := client.GetStreams(twitch.GetStreamInput{
		GameID: "33214",
		Limit:  121,
	})

	if err != nil {
		t.Errorf("Error returned from GetStreams function. Error: %v", err)
	}

	if len(s) != 121 {
		t.Errorf("Expexted 121 streams to be returned. Got %v streams", len(s))
	}
}
