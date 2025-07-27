package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/MatTwix/RE-minder-Bots/database"
	"github.com/MatTwix/RE-minder-Bots/models"
	"github.com/jackc/pgx/v5"
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

func SetChat(ctx context.Context, reminderUserID int, platform string, chatID int64) (models.Chat, error) {
	chat := models.Chat{
		ReminderUserID: reminderUserID,
		Platform:       platform,
		ChatID:         chatID,
	}

	err := database.DB.QueryRow(ctx,
		`UPDATE chats
		SET chat_id = $1
		WHERE reminder_user_id = $2 AND platform = $3
		RETURNING id, created_at, updated_at`,
		chatID, reminderUserID, platform).Scan(&chat.ID, &chat.CreatedAt, &chat.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := database.DB.QueryRow(ctx,
				`INSERT INTO chats
				(reminder_user_id, platform, chat_id)
				VALUES
				($1, $2, $3)
				RETURNING id, created_at, updated_at`,
				reminderUserID, platform, chatID).Scan(&chat.ID, &chat.CreatedAt, &chat.UpdatedAt)

			if err != nil {
				return chat, errors.New("Error creating chat: " + err.Error())
			}
		} else {
			return chat, errors.New("Error updating chat chat: " + err.Error())
		}
	}

	return chat, nil
}
