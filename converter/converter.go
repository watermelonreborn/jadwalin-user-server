package converter

import (
	"jadwalin/models"
	"time"
)

func UserRegisterToUser(userRegister models.UserRegister, authId string) models.User {
	return models.User{
		DiscordID: userRegister.DiscordID,
		ServerID:  userRegister.ServerID,
		AuthID:    authId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
