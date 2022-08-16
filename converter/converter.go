package converter

import (
	"jadwalin/models"
	"time"
)

func UserRegisterToUser(userRegister models.UserRegister) models.User {
	return models.User{
		DiscordID: userRegister.DiscordID,
		ServerID:  userRegister.ServerID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
