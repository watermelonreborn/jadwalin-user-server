package constants

import "time"

const (
	IsAuthenticatedKey     = "is_authenticated"
	UserIDKey              = "user_id"
	CodeExpirationDuration = time.Duration(60) * time.Second
)
