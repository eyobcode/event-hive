package main

import (
	"log"
	"os"

	"github.com/eyobcode/event-hive/internal/database"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must provide 'up' or 'down'")
	}

	direction := os.Args[1]
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Could not connect to DB:", err)
	}
	defer db.Close()

	instance, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fSrc, err := (&file.File{}).Open("cmd/migrate/migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, "postgres", instance)
	if err != nil {
		log.Fatal(err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("✅ Migrations applied")

	case "down":
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
		log.Println("⬇️  Migration rolled back")

	default:
		log.Fatal("Unknown direction: use 'up' or 'down'")
	}
}
