package models

type UserDiscordRequest struct {
	UserID   string `json:"user_id"`
	ServerID string `json:"server_id"`
	Token    string `json:"token"`
}
