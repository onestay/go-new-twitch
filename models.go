package twitch

import "time"

type twitchError struct {
	ErrorCode string `json:"error"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
}

func (te *twitchError) Error() string {
	return te.Message
}

type twitchResponseUser struct {
	Data []User `json:"data"`
}

// User contains information about one specificed twitch user as defined by the twitch API documentation
// https://dev.twitch.tv/docs/api/reference/#get-users
type User struct {
	// User’s ID
	ID string `json:"id"`
	// User’s login name
	Login string `json:"login"`
	// User’s display name
	DisplayName string `json:"display_name"`
	// User’s type: "staff", "admin", "global_mod", or "".
	Type string `json:"type"`
	// User’s broadcaster type: "partner", "affiliate", or "".
	BroadcasterType string `json:"broadcaster_type"`
	// User’s channel description.
	Description string `json:"description"`
	// URL of the user’s profile image.
	ProfileImageURL string `json:"profile_image_url"`
	// URL of the user’s offline image.
	OfflineImageURL string `json:"offline_image_url"`
	// Total number of views of the user’s channel.
	ViewCount int `json:"view_count"`
	// User's Email
	Email string `json:"email"`
}

type twitchResponseStream struct {
	Data       []Stream `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

// Stream contains information about one specificed stream as defined by the twitch API documentation
// https://dev.twitch.tv/docs/api/reference/#get-streams
type Stream struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	UserName     string    `json:"user_name"`
	GameID       string    `json:"game_id"`
	CommunityIds []string  `json:"community_ids"`
	Type         string    `json:"type"`
	Title        string    `json:"title"`
	ViewerCount  int       `json:"viewer_count"`
	StartedAt    time.Time `json:"started_at"`
	Language     string    `json:"language"`
	ThumbnailURL string    `json:"thumbnail_url"`
}
