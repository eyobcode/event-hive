package handlers

import (
	"log"
	"net/http"

	"github.com/eyobcode/event-hive/internal/dto"
	"github.com/eyobcode/event-hive/internal/models"
	"github.com/eyobcode/event-hive/internal/utils"
	"github.com/eyobcode/event-hive/internal/validators"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	EventModel     *models.EventModel
	OrganizerModel *models.OrganizerModel
	CategoryModel  *models.CategoryModel
}

// CreateEvent handles POST /api/events/
func (h *EventHandler) CreateEvent(c *gin.Context) {
	var input dto.EventInput

	userId, ok := utils.GetCurrentUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
		return
	}

	organizerId, err := h.OrganizerModel.EnsureOrganizer(userId)
	if err != nil {
		log.Println("Organizer check error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	if organizerId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please create the organizer account or fill the form."})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidateEventInput(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existCategory, err := h.CategoryModel.CategoryExist(*input.CategoryId)
	if err != nil {
		log.Println("Category check error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	if !existCategory {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category doesn’t exist."})
		return
	}

	event := &models.Event{
		Title:       *input.Title,
		Description: *input.Description,
		Location:    *input.Location,
		OrganizerId: organizerId,
		CategoryId:  *input.CategoryId,
		StartTime:   *input.StartTime,
		EndTime:     *input.EndTime,
	}

	if err := h.EventModel.Insert(event); err != nil {
		log.Println("DB insert error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully.",
		"event":   event,
	})
}

// GetAllEvents DetAllEvent handles GET /api/events/
func (h *EventHandler) GetAllEvents(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Get all event successfully!",
	})
}
