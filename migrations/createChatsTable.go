package migrations

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateChatsTable(DB *pgxpool.Pool) {
	ctx := context.Background()
	tx, err := DB.Begin(ctx)
	if err != nil {
		panic("Failed to begin transaction: " + err.Error())
	}
	defer tx.Rollback(ctx)

	var tableExists bool
	err = tx.QueryRow(ctx,
		"SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'chats')").
		Scan(&tableExists)
	if err != nil {
		panic("Failed to check if table exists: " + err.Error())
	}

	if !tableExists {
		_, err = tx.Exec(ctx, `
			CREATE TABLE chats (
				id SERIAL PRIMARY KEY,
				reminder_user_id BIGINT NOT NULL,
				platform VARCHAR(50) NOT NULL CHECK (platform IN ('telegram', 'discord', 'google')),
				chat_id VARCHAR(255) NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT NOW(),
				updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
				CONSTRAINT unique_chat UNIQUE (reminder_user_id, platform)
			);
		`)
		if err != nil {
			panic("Failed to create chats table: " + err.Error())
		}

		err = tx.Commit(ctx)
		if err != nil {
			panic("Failed to commit transaction: " + err.Error())
		}

		log.Println("Chats table created successfully")
	} else {
		tx.Rollback(ctx)
	}
}
