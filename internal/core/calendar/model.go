package calendar

import "time"

// SyncRequest represents the payload for syncing calendars.
type SyncRequest struct {
	Provider string `json:"provider"` // e.g., "google", "outlook"
	Token    string `json:"token"`
}

// Event represents a single calendar event.
type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Location    string    `json:"location"`
}

// EventsResponse represents the list of calendar events.
type EventsResponse struct {
	Events []Event `json:"events"`
}
