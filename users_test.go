package twitch_test

import (
	"os"
	"testing"

	twitch "github.com/onestay/go-new-twitch"
)

var clientID string

func init() {
	clientID = os.Getenv("CLIENT_ID")
}

func TestGetUsers(t *testing.T) {
	client := twitch.NewClient(clientID)

	u, err := client.GetUsersByLogin("lirik")
	if err != nil {
		t.Errorf("An error occured while getting the twitch user %v", err)
	} else if len(u) < 0 {
		t.Error("Expected one user to be returned")
	} else if u[0].Login != "lirik" {
		t.Error("Expected login name to be \"lirik\"")
	}

	tmp := make([]string, 101)
	for i := 0; i < 101; i++ {
		tmp[i] = "a"
	}

	u, err = client.GetUsersByID(tmp...)
	if err == nil {
		t.Error("Expected GetUsers to fail with 100 users")
	}
}
