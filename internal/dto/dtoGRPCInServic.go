package dto

import "time"

type RequestCreategRPCInServic struct {
	Name             string          `json:"name"`
	UserID           int             `json:"user_id"`
	Description      string          `json:"description,omitempty"`
	Timestamp        time.Time       `json:"timeStamp"`
	NotificationType map[string]bool `json:"notification_type"`
}

type ResponseCreategRPCInServic struct {
	ReminderID int `json:"reminder_id"`
}

type RequestGetgRPCInServic struct {
	ReminderID int `json:"reminder_id"`
}

type ResponseGetgRPCInServic struct {
	Name             string          `json:"name"`
	Description      string          `json:"description,omitempty"`
	Timestamp        time.Time       `json:"timeStamp"`
	NotificationType map[string]bool `json:"notification_type"`
}

type RequestUpdateGRPCInServic struct {
	Name             string          `json:"name"`
	ReminderID       int             `json:"reminder_id"`
	UserID           int             `json:"user_id"`
	Description      string          `json:"description,omitempty"`
	TimeStamp        time.Time       `json:"timeStamp"`
	NotificationType map[string]bool `json:"notification_type"`
}

type ResponseUpdateGRPCInServic struct {
	Name             string          `json:"name"`
	ReminderID       int             `json:"reminder_id"`
	Description      string          `json:"description,omitempty"`
	TimeStamp        time.Time       `json:"timeStamp"`
	NotificationType map[string]bool `json:"notification_type"`
}

type RequestDeletegRRPCInServic struct {
	ReminderID int `json:"reminder_id"`
	UserID     int `json:"user_id"`
}

func GRPCInServicMapInSliceString(nt map[string]bool) []string {

	result := make([]string, 6, 12)
	for key, value := range nt {
		if value {
			result = append(result, key)
		}
	}

	return result
}

func GRPCInServicSliceStringInMap(nt []string) map[string]bool {

	result := map[string]bool{
		"email": false,
		"tg":    false,
	}

	for _, value := range nt {

		if value == "email" {
			result[value] = true
		}

		if value == "tg" {
			result[value] = true
		}
	}

	return result
}
