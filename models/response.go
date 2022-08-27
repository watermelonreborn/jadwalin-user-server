package models

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

type UserEventsResult struct {
	Data []ReminderInput `json:"data,omitempty"`
}
