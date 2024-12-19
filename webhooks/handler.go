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

	// Process the webhook payload
	if err := w.processWebhookPayload(payload); err != nil {
		log.Printf("Error processing webhook: %v", err)
		http.Error(wr, "Failed to process webhook", http.StatusInternalServerError)
		return
	}

	// Respond to acknowledge receipt
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte("Webhook processed successfully"))
}

func (w *WebhookHandler) processWebhookPayload(payload map[string]interface{}) error {
	eventType, ok := payload["event_type"].(string)
	if !ok {
		return errors.WrapError(errors.ErrInvalidInput, "missing event_type")
	}

	log.Printf("Processing event type: %s", eventType)

	switch eventType {
	case "meeting.updated":
		// Handle meeting update
		log.Printf("Meeting updated: %v", payload)
	case "meeting.canceled":
		// Handle meeting cancellation
		log.Printf("Meeting canceled: %v", payload)
	case "calendar.sync":
		// Handle calendar synchronization
		log.Printf("Calendar sync triggered: %v", payload)
	default:
		log.Printf("Unhandled event type: %s", eventType)
		return errors.ErrNotFound
	}

	return nil
}
