package http

import (
	"net/http"

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

	// User Routes
	router.HandleFunc("/users/register", userHandler.Register).Methods(http.MethodPost)
	router.HandleFunc("/users/login", userHandler.Login).Methods(http.MethodPost)

	return router
}
