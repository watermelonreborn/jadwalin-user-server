package services

import (
	"bytes"
	"encoding/json"
	"io"
	"jadwalin/config"
	"jadwalin/models"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

func httpCall(method string, url string, body io.Reader) (*http.Response, error) {
	t := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 60 * time.Second,
	}
	client := &http.Client{Transport: t}
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

func SendReminder(reminders []models.ReminderOutput) {
	body, err := json.Marshal(reminders)

	if err != nil {
		log.Printf("[ERROR] Error creating notification/reminder json body: %s", err)
	}

	httpCall("POST", config.AppConfig.BotURL+"/reminder", bytes.NewBuffer(body))
}

func GetEventsInHour(hour int) []models.ReminderInput {
	var result models.UserEventsResult

	response, err := httpCall("GET", config.AppConfig.AuthURL+"/events/"+strconv.Itoa(hour), nil)
	if err != nil || response.StatusCode != 200 {
		log.Printf("[ERROR] Failed getting events in N hour: %v", err)
		return nil
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Printf("[ERROR] Failed decode to struct in get notification events: %v", err)
		return nil
	}

	return result.Data
}

func SendDailyNotification() {
	input := GetEventsInHour(24)

	if input == nil {
		log.Printf("[ERROR] Failed to send notification: no events found")
		return
	}

	var output []models.ReminderOutput

	// TODO: Optimize from linear to constant database calls
	for _, raw := range input {
		user, err := GetUserByAuthID(raw.UserID)
		if err != nil {
			continue
		}

		reminder := models.ReminderOutput{
			DiscordID: user.DiscordID,
			ServerID:  user.ServerID,
			Hours:     24,
			Events:    raw.Events,
		}

		output = append(output, reminder)
	}

	SendReminder(output)
}
