package models

type UserRequest struct {
	DiscordUsername string `json:"user_id"`
	GoogleAccount   string `json:"google_account"`
}
