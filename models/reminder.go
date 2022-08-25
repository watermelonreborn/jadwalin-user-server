package models

type ReminderInput struct {
	UserID string  `json:"user_id" binding:"required"`
	Events []Event `json:"events" binding:"required"`
}

type ReminderOutput struct {
	DiscordID string  `json:"discord_id" binding:"required"`
	ServerID  string  `json:"server_id" binding:"required"`
	Hours     int     `json:"hours" binding:"required"`
	Events    []Event `json:"events" binding:"required"`
}

type Event struct {
	Summary     string `json:"summary" binding:"required"`
	Description string `json:"description" binding:"required"`
	URI         string `json:"uri" binding:"required"`
}
