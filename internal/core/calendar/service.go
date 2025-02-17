package calendar

import (
	"context"
	"errors"
	"time"
)

// CalendarService provides business logic for calendar operations.
type CalendarService struct {
	repo *MongoDBCalendarRepository
}

// NewCalendarService creates a new instance of CalendarService.
func NewCalendarService(repo *MongoDBCalendarRepository) *CalendarService {
	return &CalendarService{repo: repo}
}

// SyncCalendar handles synchronization with external calendar providers.
func (s *CalendarService) SyncCalendar(ctx context.Context, request SyncRequest) error {
	// Validate provider input
	if request.Provider == "" || request.Token == "" {
		return errors.New("invalid provider or token")
	}

	// Example: Perform specific provider sync logic (expand as needed)
	if request.Provider == "google" {
		// Google calendar sync logic
		// Example: Call external API to fetch events
	} else if request.Provider == "outlook" {
		// Outlook calendar sync logic
	} else {
		return errors.New("unsupported provider")
	}

	// Save events to the repository (stub for now)
	events := []Event{
		{ID: "1", Title: "Meeting 1", StartTime: time.Now(), EndTime: time.Now().Add(1 * time.Hour)},
	}
	return s.repo.SaveEvents(ctx, events)
}

// GetEvents fetches all calendar events from the repository.
func (s *CalendarService) GetEvents(ctx context.Context) ([]Event, error) {
	events, err := s.repo.GetAllEvents(ctx)
	if err != nil {
		return nil, errors.New("failed to retrieve events")
	}
	return events, nil
}
