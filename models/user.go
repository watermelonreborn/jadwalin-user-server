package models

import "time"

type User struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	GoogleAccount string    `json:"google_account" gorm:"not null"`
	DiscordID     string    `json:"discord_id" gorm:"not null"`
	ServerID      string    `json:"-" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UserRegister struct {
	DiscordID string `json:"discord_id" binding:"required"`
	ServerID  string `json:"server_id" binding:"required"`
	Code      string `json:"code" binding:"required"`
}
