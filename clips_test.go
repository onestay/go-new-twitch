package twitch_test

import (
	"fmt"
	"testing"

	twitch "github.com/onestay/go-new-twitch"
)

func TestGetClips(t *testing.T) {
	i := twitch.GetClipsInput{
		BroadcasterID: "27446517",
	}

	cd, err := client.GetClips(i)
	fmt.Println(cd)
	if err != nil {
		t.Errorf("An error occurred while getting the clips: %v", err)
	} else if len(cd) != 20 {
		t.Errorf("expected 20 clips being returned")
	} else if cd[0].BroadcasterName != "monstercat" {
		t.Errorf("expected broadcaster_name to be monstercat")
	}
}
