package application

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func SendMessageToDiscord(msg string) int {
	discordWebhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if discordWebhookURL == "" {
		log.Print("Error: el webhook de Discord no está configurado")
		return http.StatusInternalServerError
	}

	payload := map[string]string{"content": msg}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error al convertir mensaje a JSON: %v", err)
		return http.StatusInternalServerError
	}

	resp, err := http.Post(discordWebhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("Error al enviar mensaje a Discord: %v", err)
		return http.StatusInternalServerError
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
