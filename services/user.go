package services

import (
	"fmt"
	"jadwalin/models"
)

func CreateUser(userRegister models.UserRegister) (string, error) {
	// TODO: create logic to write user to database and save user's server
	fmt.Println(userRegister)
}

func GetUser(userId string) {
	// TODO: create logic to get user from database
	fmt.Println(userId)
}

func DeleteUser(userId string) {
	// TODO: create logic to delete user from database
	fmt.Println(userId)
}
