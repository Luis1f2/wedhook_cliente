package application

import (
	"encoding/json"
	"fmt"
	domain "github_wb/domain/value_objects"
	"log"
)

type StarEventPayload struct {
	Sender struct {
		Login   string `json:"login"`
		HTMLURL string `json:"html_url"`
	} `json:"sender"`
	Repository struct {
		FullName string `json:"full_name"`
		HTMLURL  string `json:"html_url"`
	} `json:"repository"`
}

func ProcessPullRequest(payload []byte) string {
	var eventPayload domain.PullRequestEventPayload

	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		log.Printf("Error al deserializar el payload: %v", err)
		return "ERROR"
	}

	if eventPayload.Action == "closed" {
		base := eventPayload.PullRequest.Base.Ref
		branch := eventPayload.PullRequest.Head.Ref
		user := eventPayload.PullRequest.User.Login
		title := eventPayload.PullRequest.Title
		url := eventPayload.PullRequest.URL

		message := fmt.Sprintf("🔔 **Pull Request Cerrado** 🔔\n👤 Autor: %s\n📌 Título: %s\n🔀 De: %s ➝ %s\n🔗 [Ver Pull Request](%s)", user, title, branch, base, url)

		log.Printf("Notificación generada para Discord: %s", message)
		return message
	}

	log.Printf("Pull Request no cerrado: %s", eventPayload.Action)
	return "ERROR"
}

func ProcessStarEvent(payload []byte) string {
	var eventPayload StarEventPayload

	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		log.Printf("Error al deserializar el payload de star: %v", err)
		return "ERROR"
	}

	message := fmt.Sprintf("⭐ ¡Nuevo star en el repositorio! ⭐\n👤 Usuario: [%s](%s)\n📂 Repositorio: [%s](%s)",
		eventPayload.Sender.Login, eventPayload.Sender.HTMLURL,
		eventPayload.Repository.FullName, eventPayload.Repository.HTMLURL)

	log.Printf("Notificación generada para Discord: %s", message)
	return message
}
