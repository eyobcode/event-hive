package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	DB *sql.DB
}

// CreateEvent handles POST /api/events/
func (h *EventHandler) CreateEvent(c *gin.Context) {
	// For now, just a placeholder
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully!",
	})
}

// GetAllEvents DetAllEvent handles GET /api/events/
func (h *EventHandler) GetAllEvents(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Get all event successfully!",
	})
}
