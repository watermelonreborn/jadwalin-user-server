package models

import (
	"time"
)

type Server struct {
	ID        string                 `json:"server_id" bson:"_id,omitempty"`
	Members   map[string]interface{} `json:"members"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}
