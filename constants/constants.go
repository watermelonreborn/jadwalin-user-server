package constants

import "time"

const (
	IsAuthenticatedKey     = "is_authenticated"
	UserIDKey              = "user_id"
	CodeExpirationDuration = time.Duration(60) * time.Second

	AlreadyRegistered = "ALREADY_REGISTERED"
	Registered        = "REGISTERED"
	Success           = "SUCCESS"
	Error             = "ERROR"

	LogInfo  = "[INFO]"
	LogError = "[ERROR]"
)
