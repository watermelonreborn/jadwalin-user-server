package models

import (
	"time"
)

type Server struct {
	ID        string    `json:"server_id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Members   []User    `json:"members" gorm:"foreignKey:ServerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
