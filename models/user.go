package models

import "time"

type User struct {
	ID              string    `json:"user_id"`
	DiscordUsername string    `json:"discord_username"`
	GoogleAccount   string    `json:"google_account"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
