package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eyobcode/event-hive/internal/config"
)

func ConnectDB() *sql.DB {
	cfg := config.LoadConfig()

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
		log.Fatal("❌ Failed to connect to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("❌ Database unreachable:", err)
	}

	log.Println("✅ Connected to PostgreSQL")
	return db
}
