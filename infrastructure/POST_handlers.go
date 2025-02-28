package infrastructure

import (
	"github_wb/application"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PullRequestEvent(ctx *gin.Context) {
	eventType := ctx.GetHeader("X-GitHub-Event")
	deliveryID := ctx.GetHeader("X-GitHub-Delivery")
	signature := ctx.GetHeader("X-Hub-Signature-256")

	log.Println(signature)

	log.Printf("Webhook recibido: \nEvento=%s, \nDeliveryID=%s", eventType, deliveryID)

	payload, err := ctx.GetRawData()
	if err != nil {
		log.Printf("Error al leer el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}

	var statusCode int
	var msg string

	switch eventType {
	case "pull_request":
		msg = application.ProcessPullRequest(payload)
	case "star":
		msg = application.ProcessStarEvent(payload)
	default:
		ctx.JSON(http.StatusOK, gin.H{"status": "Evento no soportado"})
		return
	}

	if msg == "ERROR" {
		statusCode = http.StatusInternalServerError
	} else {
		statusCode = application.SendMessageToDiscord(msg)
	}

	switch statusCode {
	case http.StatusOK:
		ctx.JSON(http.StatusOK, gin.H{"status": "Evento recibido y procesado"})
	case http.StatusInternalServerError:
		log.Printf("Error al procesar el evento")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el evento"})
	default:
		ctx.JSON(http.StatusOK, gin.H{"status": "Peticion procesada"})
	}
}
