package dto

import "time"

// dto Data Transfer Object

type EventInput struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Location    *string    `json:"location"`
	CategoryId  *int       `json:"category_id"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
}
