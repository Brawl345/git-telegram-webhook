package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/Brawl345/gitwebhook/webhook"
	webhookEvent "github.com/Brawl345/gitwebhook/webhook/event"
)

const (
	TelegramAPIURL = "https://api.telegram.org/bot%s/sendMessage"
)

func SendTelegramMessage(token, chatID, message string) error {
	apiURL := fmt.Sprintf(TelegramAPIURL, token)
	resp, err := http.PostForm(apiURL, url.Values{
		"chat_id":                  {chatID},
		"text":                     {message},
		"parse_mode":               {"HTML"},
		"disable_web_page_preview": {"true"},
	})
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram API error: %s", body)
	}
	return nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	chatID := r.URL.Query().Get("chat_id")
	if chatID == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println("Missing chat_id in query parameters")
		return
	}

	telegramToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if telegramToken == "" {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Missing TELEGRAM_BOT_TOKEN")
		return
	}

	secret := os.Getenv("GITHUB_WEBHOOK_SECRET")
	if secret == "" {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Missing GITHUB_WEBHOOK_SECRET")
		return
	}

	signature := r.Header.Get("X-Hub-Signature-256")
	if signature == "" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		log.Println("Missing X-Hub-Signature-256")
		return
	}

	event := r.Header.Get("X-GitHub-Event")
	if event == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println("Missing X-GitHub-Event")
		return
	}

	var data webhookEvent.Event

	switch event {
	case webhookEvent.PushEventName:
		data = &webhookEvent.PushPayload{}
		break
	case webhookEvent.PingEventName:
		data = &webhookEvent.PingPayload{}
		break
	default:
		http.Error(w, "Not Implemented", http.StatusNotImplemented)
		log.Println("Unsupported X-GitHub-Event:", event)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println("Error reading request body:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(r.Body)

	if os.Getenv("DANGEROUS_SKIP_GITHUB_WEBHOOK_SECRET_CHECK") == "" {
		if !webhook.VerifySignature(body, secret, signature) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			log.Println("Invalid signature")
			return
		}
	} else {
		log.Println("Skipping signature check, this is DANGEROUS!")
	}

	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	if err := SendTelegramMessage(telegramToken, chatID, data.Text()); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error sending message to Telegram:", err)
		return
	}

	_, err = fmt.Fprintf(w, "Webhook handled successfully")
	if err != nil {
		log.Println("Error writing response:", err)
		return
	}
}
