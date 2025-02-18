package handlers

import (
	"encoding/json"
	"net/http"

	"meetup-app-hexa-arch/internal/core/meeting"
)

type MeetingHandler struct {
	meetingService *meeting.MeetingService
}

func NewMeetingHandler(meetingService *meeting.MeetingService) *MeetingHandler {
	return &MeetingHandler{meetingService: meetingService}
}

func (h *MeetingHandler) ScheduleMeeting(w http.ResponseWriter, r *http.Request) {
	var req meeting.Meeting
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.meetingService.ScheduleMeeting(req)
	if err != nil {
		http.Error(w, "Failed to schedule meeting", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Meeting scheduled successfully"})
}

func (h *MeetingHandler) GetMeeting(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing meeting ID", http.StatusBadRequest)
		return
	}

	meeting, err := h.meetingService.GetMeetingByID(id)
	if err != nil {
		http.Error(w, "Meeting not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meeting)
}
