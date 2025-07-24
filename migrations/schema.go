package migrations

import "github.com/jackc/pgx/v5/pgxpool"

func Migrate(DB *pgxpool.Pool) {
	CreateChatsTable(DB)
}
