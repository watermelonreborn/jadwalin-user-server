package models

import (
	"time"
)

type Server struct {
	ID          string                 `json:"server_id" bson:"_id,omitempty"`
	TextChannel string                 `json:"text_channel" bson:"text_channel"`
	Members     map[string]interface{} `json:"members"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

type ServerCreate struct {
	ServerID    string `json:"server_id" binding:"required"`
	TextChannel string `json:"text_channel" binding:"required"`
}
