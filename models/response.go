package models

type Response struct {
	Code  int         `json:"code,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

type UserEventsResult struct {
	Data []ReminderInput `json:"data,omitempty"`
}
