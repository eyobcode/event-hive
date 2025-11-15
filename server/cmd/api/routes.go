package main

import (
	"net/http"

	"github.com/eyobcode/event-hive/internal/handlers"
	"github.com/eyobcode/event-hive/internal/models"
	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	router := gin.Default()
	m := models.NewModels(app.DB)

	api := router.Group("/api")
	{
		event := api.Group("/events")
		{
			h := handlers.EventHandler{NewModels: &m}
			event.POST("/", h.CreateEvent)
			event.GET("/", h.GetAllEvents)
		}

	}

	return router
}
