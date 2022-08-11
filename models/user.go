package models

type User struct {
	ID            string `json:"user_id" gorm:"type:varchar(30);unique;primary_key"`
	GoogleAccount string `json:"google_account" gorm:"not null"`
}
