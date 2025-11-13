package main

import (
	"net/http"

	"github.com/eyobcode/event-hive/internal/handlers"
	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	router := gin.Default()
	api := router.Group("/api")
	{
		event := api.Group("/events")
		{
			h := handlers.EventHandler{DB: app.DB}
			event.POST("/", h.CreateEvent)
			event.GET("/", h.GetAllEvents)
		}
	}

	return router
}
