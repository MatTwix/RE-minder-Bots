package models

import "time"

type Chat struct {
	ID             int64     `json:"id"`
	ReminderUserID int64     `json:"reminder_user_id"`
	Platform       string    `json:"platform"`
	ChatID         int64     `json:"chat_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
