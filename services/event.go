package services

import (
	"encoding/json"
	"jadwalin/config"
	"jadwalin/models"
	"log"
	"strconv"

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

// TODO: Implement
func SendReminder(reminders []models.ReminderOutput) (*http.Response, error) {
	return nil, nil
}

func GetEventsInHour(hour int) {
	var (
		result models.UserEventsResult
	)

	response, err := httpCall("GET", config.AppConfig.AuthURL+"/events/"+strconv.Itoa(hour))
	if response.StatusCode != 200 || err != nil {
		log.Printf("Error getting events in N hour: %v", err)
		return
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Printf("Error decode to struct: %v", err)
		return
	}

	//TODO: process the result

}
