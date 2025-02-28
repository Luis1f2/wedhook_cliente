package application

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func send_message_to_discord(msg string) int {
	discord_webhook_url := os.Getenv("DISCORD_WEBHOOK_URL")
	if discord_webhook_url == "" {
		log.Print("Error: El webhook de Discord no está configurado")
		return 500
	}

	payload := map[string]string{"content": msg}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error al convertir mensaje a JSON: %v", err)
		return 500
	}

	resp, err := http.Post(discord_webhook_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("Error al enviar mensaje a Discord: %v", err)
		return 500
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
