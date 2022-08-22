package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DiscordID string             `json:"discord_id" bson:"discord_id"`
	ServerID  string             `json:"server_id" bson:"server_id"`
	AuthID    string             `json:"-" bson:"auth_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type UserRegister struct {
	DiscordID string `json:"discord_id" binding:"required"`
	ServerID  string `json:"server_id" binding:"required"`
	Code      string `json:"code" binding:"required"`
}

type UserResponse struct {
	DiscordID          string `json:"discord_id"`
	ServerID           string `json:"server_id"`
	RegistrationStatus string `json:"registration_status"`
}
