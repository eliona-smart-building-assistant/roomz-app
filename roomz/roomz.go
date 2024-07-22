package roomz

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type WebhookEvent struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Timestamp  string                 `json:"timestamp"`
	ApiVersion string                 `json:"apiVersion"`
	Data       map[string]interface{} `json:"data"`
}

type PresenceStatus string

const (
	Free PresenceStatus = "free"
	Busy PresenceStatus = "busy"
)

type WorkspacePresenceChangedEvent struct {
	WorkspaceId    string         `json:"workspaceId"`
	PresenceStatus PresenceStatus `json:"presenceStatus"`
}

type WebhookHandlerFunc func(WebhookEvent) error

type WebhookServer struct {
	handlers map[string]WebhookHandlerFunc
}

func NewWebhookServer() *WebhookServer {
	return &WebhookServer{
		handlers: make(map[string]WebhookHandlerFunc),
	}
}

func (s *WebhookServer) RegisterHandler(eventType string, handler WebhookHandlerFunc) {
	s.handlers[eventType] = handler
}

func (s *WebhookServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	signature := r.Header.Get("X-Roomz-Signature")
	if signature == "" {
		http.Error(w, "Missing X-Roomz-Signature header", http.StatusBadRequest)
		return
	}

	// todo: Validate signature?

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var event WebhookEvent
	if err := json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	handler, ok := s.handlers[event.Type]
	if !ok {
		http.Error(w, "Unsupported event type", http.StatusBadRequest)
		return
	}

	if err := handler(event); err != nil {
		http.Error(w, "Failed to handle event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func HandleWorkspacePresenceChanged(callback func(workspaceId string, presenceStatus PresenceStatus) error) WebhookHandlerFunc {
	return func(event WebhookEvent) error {
		data := event.Data

		workspaceId, ok := data["workspaceId"].(string)
		if !ok {
			return errors.New("invalid workspaceId format")
		}

		presenceStatus, ok := data["presenceStatus"].(string)
		if !ok {
			return errors.New("invalid presenceStatus format")
		}

		return callback(workspaceId, PresenceStatus(presenceStatus))
	}
}
