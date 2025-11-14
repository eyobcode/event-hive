package validators

import (
	"fmt"
	"strings"
	"time"

	"github.com/eyobcode/event-hive/internal/dto"
)

func ValidateEventInput(input *dto.EventInput) error {
	missing := []string{}
	now := time.Now()

	if input.Title == nil || strings.TrimSpace(*input.Title) == "" {
		missing = append(missing, "title")
	}
	if input.Description == nil || strings.TrimSpace(*input.Description) == "" {
		missing = append(missing, "description")
	}
	if input.CategoryId == nil || *input.CategoryId == 0 {
		missing = append(missing, "category id")
	}
	if input.Location == nil || strings.TrimSpace(*input.Location) == "" {
		missing = append(missing, "location")
	}
	if input.StartTime == nil || input.StartTime.Before(now) {
		missing = append(missing, "start time must be in the future")
	}
	if input.EndTime == nil || input.EndTime.Before(now) {
		missing = append(missing, "end time must be in the future")
	}
	if input.StartTime != nil && input.EndTime != nil && input.EndTime.Before(*input.StartTime) {
		missing = append(missing, "end time cannot be before start time")
	}

	if len(missing) > 0 {
		return fmt.Errorf("validation errors: %s", strings.Join(missing, " \n "))
	}

	return nil
}
