package http

import (
	"meetup-app-hexa-arch/internal/adapters/auth"
	"meetup-app-hexa-arch/internal/adapters/http/handlers"
	"meetup-app-hexa-arch/internal/adapters/http/middleware"
	"meetup-app-hexa-arch/webhooks"

	"github.com/gorilla/mux"
)

func NewRouter(userHandler *handlers.UserHandler, meetingHandler *handlers.MeetingHandler, calendarHandler *handlers.CalendarHandler, jwtService *auth.JWTService, webhookHandler *webhooks.WebhookHandler) *mux.Router {
	router := mux.NewRouter()

	// Middleware
	authMiddleware := middleware.NewAuthMiddleware(jwtService)
	router.Use(authMiddleware.AuthMiddleware)

	return router
}
