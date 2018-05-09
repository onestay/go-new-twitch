package twitch_test

import (
	"testing"
)

func TestGetGames(t *testing.T) {
	g, err := client.GetGameByName("overwatch", "League of Legends")
	if err != nil {
		t.Errorf("An error occured while getting the games: %v", err)
		return
	}
	if len(g) != 2 {
		t.Errorf("Expected 2 games to be returned. Got %v", len(g))
	}
	if g[0].Name != "overwatch" {
		t.Errorf("Expected first game to be returned to be overwatch. Got %v", g[0].Name)
	}
}
