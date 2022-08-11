package models

import "time"

type Server struct {
	ID        string    `json:"server_id"`
	Name      string    `json:"name"`
	Members   []User    `json:"members"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
