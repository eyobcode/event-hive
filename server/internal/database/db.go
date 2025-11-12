package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eyobcode/event-hive/internal/config"
	_ "github.com/lib/pq"
)

// ConnectDB loads config internally and connects to PostgreSQL
func ConnectDB() (*sql.DB, error) {
	cfg := config.LoadConfig() // load database config internally

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Pass,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database unreachable: %w", err)
	}

	log.Println("Connected to PostgreSQL....")
	return db, nil
}
