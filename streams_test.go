package twitch_test

import (
	"fmt"
	"testing"

	twitch "github.com/onestay/go-new-twitch"
)

func TestGetStreams(t *testing.T) {
	// it's kinda hard to test this, since it requires the channel to be online. So this test might be a bit unreliable.
	i := twitch.GetStreamsInput{
		UserLogin: []string{"monstercat"},
	}

	sd, err := client.GetStreams(i)
	fmt.Println(sd)
	if err != nil {
		t.Errorf("An error occured while getting the streams: %v", err)
	} else if len(sd) != 1 {
		t.Errorf("expected one stream being returned")
	} else if sd[0].UserID != "27446517" {
		t.Errorf("expected user_id to be 27446517")
	}
}
