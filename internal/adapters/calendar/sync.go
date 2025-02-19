package calendar

type CalendarSyncService struct {
	Google  *GoogleCalendarService
	Outlook *OutlookCalendarService
}

func NewCalendarSyncService(google *GoogleCalendarService, outlook *OutlookCalendarService) *CalendarSyncService {
	return &CalendarSyncService{
		Google:  google,
		Outlook: outlook,
	}
}

func (c *CalendarSyncService) Sync() {
	c.Google.Sync()
	c.Outlook.Sync()
}
