package services

import (
	"jadwalin/config"
	"net/http"
)

func SyncCalendar(authId string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", config.AppConfig.AuthURL+"/sync/"+authId, nil)
	res, err := client.Do(req)
	return res, err
}
