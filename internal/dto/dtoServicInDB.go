package dto

import "time"

type RequestCreateServicInDB struct {
	Name             string          `json:"name"`
	UserID           int             `json:"user_id"`
	Description      string          `json:"desription"`
	TimeStamp        time.Time       `json:"timeStamp"`
	NotificationType map[string]bool `json:"notification_type"`
}

type ResponseCreateServicInDB struct {
	ReminderID int `json:"reminder_id"`
}

type RequestGetServicInDB struct {
	ReminderID int `json:"reminder_id"`
}

type ResponseGetServicInDB struct {
	Name             string          `json:"name"`
	Description      string          `json:"desription"`
	TimeStamp        time.Time       `json:"timeStamp"`
	NotificationType map[string]bool `json:"notification_type"`
}

type RequestUpdateServicInDB struct {
	Name             string          `json:"name"`
	ReminderID       int             `json:"reminder_id"`
	UserID           int             `json:"user_id"`
	Description      string          `json:"desription"`
	TimeStamp        time.Time       `json:"timeStamp"`
	NotificationType map[string]bool `json:"notification_type"`
}

type ResponseUpdateServicInDB struct {
	Name             string          `json:"name"`
	ReminderID       int             `json:"reminder_id"`
	TimeStamp        time.Time       `json:"timeStamp"`
	Description      string          `json:"desription"`
	NotificationType map[string]bool `json:"notification_type"`
}

type RequestDeleteServicInDB struct {
	ReqeuestID int `json:"reminder_id"`
	UserID     int `json:"user_id"`
}
