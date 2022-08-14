package services

import (
	"context"

	"jadwalin/constants"
)

func CreateCode(userId string) (string, error) {
	// TODO: Check of code already exists in redis
	code := "example"

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
