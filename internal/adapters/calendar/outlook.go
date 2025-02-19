package calendar

type OutlookCalendarService struct {
	APIKey string
}

func NewOutlookCalendarService(apiKey string) *OutlookCalendarService {
	return &OutlookCalendarService{APIKey: apiKey}
}

func (o *OutlookCalendarService) Sync() {
	// Logic to sync with Outlook Calendar
}
