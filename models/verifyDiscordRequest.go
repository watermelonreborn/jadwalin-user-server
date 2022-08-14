package models

type VerifyDiscordRequest struct {
	DiscordUsername string `json:"user_id"`
	ServerID        string `json:"server_id"`
	Token           string `json:"token"`
}
