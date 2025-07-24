package database

import (
	"context"
	"log"
	"time"

	"github.com/MatTwix/RE-minder-Bots/config"
	"github.com/MatTwix/RE-minder-Bots/migrations"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	cfg := config.LoadConfig()

	if cfg.DatabaseURL == "" {
		panic("DATABASE_URL is not set in the environment variables")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		panic("Unable to connect to database: " + err.Error())
	}

	DB = pool
	log.Print("Connected to the database successfully")

	migrations.Migrate(DB)
}
