package handlers

import (
	"encoding/json"
	"meetup-app-hexa-arch/internal/core/calendar"
	"net/http"
)

type CalendarHandler struct {
	calendarService *calendar.CalendarService
}

func NewCalendarHandler(calendarService *calendar.CalendarService) *CalendarHandler {
	return &CalendarHandler{calendarService: calendarService}
}

// SyncCalendar handles calendar synchronization with external providers (Google, Outlook, etc.).
func (h *CalendarHandler) SyncCalendar(w http.ResponseWriter, r *http.Request) {
	// Parse the request payload
	var request calendar.SyncRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the service layer for syncing
	err := h.calendarService.SyncCalendar(r.Context(), request)
	if err != nil {
		http.Error(w, "Failed to sync calendar: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Calendar synced successfully"}`))
}

// ListEvents handles listing events from a connected calendar.
func (h *CalendarHandler) ListEvents(w http.ResponseWriter, r *http.Request) {
	// Retrieve the events via the service layer
	events, err := h.calendarService.GetEvents(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch calendar events: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the list of events
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
