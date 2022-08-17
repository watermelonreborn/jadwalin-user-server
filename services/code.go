package services

import (
	"context"
	"math/rand"

	"jadwalin/constants"
)

func CreateCode(userId string) (string, error) {
	// TODO: Check of code already exists in redis
	code := generateRandomString(constants.CodeLength)

	rdb := RedisClient
	res := rdb.Set(context.Background(), code, userId, constants.CodeExpirationDuration)

	return code, res.Err()
}

func UseCode(code string) (string, error) {
	// TODO: Remove code after use
	rdb := RedisClient
	res := rdb.Get(context.Background(), code)

	return res.Val(), res.Err()
}

func generateRandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
