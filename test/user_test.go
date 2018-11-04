package test

import (
	"testing"

	twitch "github.com/onestay/go-new-twitch"
)

func TestUser(t *testing.T) {
	client := twitch.NewClient("")

	u, err := client.GetUserByLogin("onestay")
	if err != nil {
		t.Errorf("Error returned from GetUserByLogin function, error: %v", err)
	}

	if u.Login != "onestay" {
		t.Errorf("GetUserByLogin incorrect, got: %v, want %v", u.Login, "onestay")
	}

	_, err = client.GetUserByLogin("")
	if err == nil {
		t.Errorf("An error should be returned on empty call")
	}

	u, err = client.GetUserByID("38782510")
	if err != nil {
		t.Errorf("Error returned from GetUserByLogin function, error: %v", err)
	}

	if u.Login != "onestay" {
		t.Errorf("GetUserByLogin incorrect, got: %v, want %v", u.Login, "onestay")
	}

	input := twitch.GetUsersInput{
		Login: []string{"onestay"},
		ID:    []string{"26301881"},
	}

	users, err := client.GetUsers(input)
	if err != nil {
		t.Errorf("Error returned from GetUserByLogin function, error: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("Expected 2 users to be returned from the API, got %v", len(users))
	}
}
