package services

import (
	"jadwalin/config"
	"net/http"
)

func httpCall(method string, url string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)
	res, err := client.Do(req)
	return res, err
}

func SyncCalendar(authId string) (*http.Response, error) {
	return httpCall("GET", config.AppConfig.AuthURL+"/sync/"+authId)
}

func GetEvents(authId string) (*http.Response, error) {
	return httpCall("GET", config.AppConfig.AuthURL+"/user-events/"+authId)
}

// TODO: Implement
func GetSummary(authId string, days int) (*http.Response, error) {
	return nil, nil
}
