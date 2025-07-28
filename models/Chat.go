package models

import "time"

type Chat struct {
	ID             int       `json:"id"`
	ReminderUserID int       `json:"reminder_user_id"`
	Platform       string    `json:"platform"`
	ChatID         string    `json:"chat_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
