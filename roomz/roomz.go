//  This file is part of the Eliona project.
//  Copyright © 2024 IoTEC AG. All Rights Reserved.
//  ______ _ _
// |  ____| (_)
// | |__  | |_  ___  _ __   __ _
// |  __| | | |/ _ \| '_ \ / _` |
// | |____| | | (_) | | | | (_| |
// |______|_|_|\___/|_| |_|\__,_|
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING
//  BUT NOT LIMITED  TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//  NON INFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
//  DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package roomz

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/eliona-smart-building-assistant/go-utils/log"
)

type webhookEvent struct {
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

type workspacePresenceChangedEvent struct {
	WorkspaceId    string         `json:"workspaceId"`
	PresenceStatus PresenceStatus `json:"presenceStatus"`
}

type webhookHandlerFunc func(workspaceId string, presenceStatus PresenceStatus) error

type webhookServer struct {
	handlers map[string]webhookHandlerFunc
	secret   string
}

func newWebhookServer(secret string) *webhookServer {
	return &webhookServer{
		handlers: make(map[string]webhookHandlerFunc),
		secret:   secret,
	}
}

func (s *webhookServer) registerHandler(eventType string, handler webhookHandlerFunc) {
	s.handlers[eventType] = handler
}

func (s *webhookServer) serveHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("webhook", "Received request for URL: %s, Method: %s", r.URL.Path, r.Method)

	signature := r.Header.Get("Roomz-Signature")
	if signature == "" {
		log.Warn("webhook", "Missing Roomz-Signature header")
		http.Error(w, "Missing X-Roomz-Signature header", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Warn("webhook", "Failed to read request body: %v", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if !s.checkSignature(body, signature) {
		log.Warn("webhook", "Invalid signature: %v", signature)
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}
	log.Debug("webhook", "Signature verified successfully")

	var event webhookEvent
	if err := json.Unmarshal(body, &event); err != nil {
		log.Warn("webhook", "Invalid JSON: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	handler, ok := s.handlers[event.Type]
	if !ok {
		log.Warn("webhook", "Unsupported event type: %s", event.Type)
		http.Error(w, "Unsupported event type", http.StatusBadRequest)
		return
	}

	if err := handler(event.Data["workspaceId"].(string), PresenceStatus(event.Data["presenceStatus"].(string))); err != nil {
		log.Warn("webhook", "Failed to handle event: %v", err)
		http.Error(w, "Failed to handle event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Debug("webhook", "Request handled successfully")
}

func (s *webhookServer) checkSignature(payload []byte, receivedSignature string) bool {
	generatedSignature := s.generateSignature(payload)
	return hmac.Equal([]byte(generatedSignature), []byte(receivedSignature))
}

func (s *webhookServer) generateSignature(payload []byte) string {
	h := hmac.New(sha512.New, []byte(s.secret))
	h.Write(payload)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func handleWorkspacePresenceChanged(callback func(workspaceId string, presenceStatus PresenceStatus) error) webhookHandlerFunc {
	return func(workspaceId string, status PresenceStatus) error {
		return callback(workspaceId, status)
	}
}

func StartWebhookListener(secret string, callback func(workspaceId string, presenceStatus PresenceStatus) error) {
	server := newWebhookServer(secret)
	server.registerHandler("workspace.presence.changed", handleWorkspacePresenceChanged(callback))
	http.HandleFunc("/webhook", server.serveHTTP)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Debug("webhook", "Received request for URL: %s, Method: %s", r.URL.Path, r.Method)
		http.Error(w, "Not Found", http.StatusNotFound)
		log.Warn("webhook", "404 Not Found for URL: %s", r.URL.Path)
	})
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("roomz webhook", "Error starting server on port 8081: %v\n", err)
	}
}
