package models

import "time"

type User struct {
	ID        uint      `json:"id" bson:"_id,omitempty"`
	DiscordID string    `json:"discord_id" bson:"discord_id"`
	ServerID  string    `json:"server_id" bson:"server_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type UserRegister struct {
	DiscordID string `json:"discord_id" binding:"required"`
	ServerID  string `json:"server_id" binding:"required"`
	Code      string `json:"code" binding:"required"`
}
