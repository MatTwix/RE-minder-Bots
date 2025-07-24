package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/MatTwix/RE-minder-Bots/database"
	"github.com/MatTwix/RE-minder-Bots/models"
)

func GetChats(ctx context.Context, optCondition ...Condition) ([]models.Chat, error) {
	whereStatement := ""
	args := []any{}
	if len(optCondition) > 0 {
		whereStatement = fmt.Sprintf("WHERE %s %s $1", optCondition[0].Field, optCondition[0].Operator)
		args = append(args, optCondition[0].Value)
	}

	var chats []models.Chat

	rows, err := database.DB.Query(ctx, fmt.Sprintf(`
		SELECT id, reminder_user_id, platform, chat_id, created_at, updated_at FROM chats %s
	`, whereStatement), args...)

	if err != nil {
		return nil, errors.New("failed to query chats: " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var chat models.Chat
		if err := rows.Scan(&chat.ID, &chat.ReminderUserID, &chat.Platform, &chat.ChatID, &chat.CreatedAt, &chat.UpdatedAt); err != nil {
			return chats, errors.New("error parsing data: " + err.Error())
		}
		chats = append(chats, chat)
	}

	return chats, nil
}
