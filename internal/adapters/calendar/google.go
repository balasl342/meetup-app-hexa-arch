package calendar

type GoogleCalendarService struct {
	APIKey string
}

func NewGoogleCalendarService(apiKey string) *GoogleCalendarService {
	return &GoogleCalendarService{APIKey: apiKey}
}

func (g *GoogleCalendarService) Sync() {
	// Logic to sync with Google Calendar
}
