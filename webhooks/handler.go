package webhooks

import (
	"encoding/json"
	"log"
	"net/http"

	"meetup-app-hexa-arch/internal/shared/errors"
)

type WebhookHandler struct{}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{}
}

// HandleWebhook processes incoming webhook requests.
func (w *WebhookHandler) HandleWebhook(wr http.ResponseWriter, r *http.Request) {
	var payload map[string]interface{}

	// Parse incoming JSON payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Failed to parse webhook payload: %v", err)
		http.Error(wr, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Printf("Webhook received: %v", payload)

}

