package models

import "time"

type User struct {
	ID            string    `json:"user_id"`
	GoogleAccount string    `json:"google_account"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
