package main

import (
	"database/sql"
	"log"

	"github.com/eyobcode/event-hive/internal/config"
	"github.com/eyobcode/event-hive/internal/database"
	"github.com/eyobcode/event-hive/internal/models"
)

type application struct {
	config config.Config
	models models.Models
	DB     *sql.DB
}

func main() {
	// 1. Load configuration (for port, JWT, etc.)
	cfg := config.LoadConfig()

	// 2. Connect to the database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Could not connect to DB:", err)
	}
	defer db.Close()

	// 3. Initialize models
	model := models.NewModels(db)

	// 4. Create application instance
	app := &application{
		config: cfg,
		models: model,
		DB:     db,
	}

	// 5. Start HTTP server
	log.Printf("🚀 Server running on port %d", app.config.Port)
	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
