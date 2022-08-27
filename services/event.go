package services

import (
	"bytes"
	"encoding/json"
	"io"
	"jadwalin/config"
	"jadwalin/models"
	"log"
	"strconv"

	"net/http"
)

func httpCall(method string, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, body)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := client.Do(req)
	return res, err
}

func SyncCalendar(authId string) (*http.Response, error) {
	return httpCall("GET", config.AppConfig.AuthURL+"/sync/"+authId, nil)
}

func GetEvents(authId string) (*http.Response, error) {
	return httpCall("GET", config.AppConfig.AuthURL+"/user-events/"+authId, nil)
}

// TODO: Implement
func GetSummary(authId string, days int, startHour int, endHour int) (*http.Response, error) {
	body, err := json.Marshal(map[string]interface{}{
		"user_id":    authId,
		"days":       days,
		"start_hour": startHour,
		"end_hour":   endHour,
	})

	if err != nil {
		log.Printf("[ERROR] Error creating summary json body: %s", err)
	}

	return httpCall("POST", config.AppConfig.AuthURL+"/user-summary", bytes.NewBuffer(body))
}

// TODO: Implement
func SendReminder(reminders []models.ReminderOutput) (*http.Response, error) {
	return nil, nil
}

func GetEventsInHour(hour int) {
	var (
		result models.UserEventsResult
	)

	response, err := httpCall("GET", config.AppConfig.AuthURL+"/events/"+strconv.Itoa(hour), nil)
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
