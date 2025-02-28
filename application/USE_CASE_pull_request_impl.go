package application

import (
	"encoding/json"
	domain "github_wb/domain/value_objects"
	"log"
)

func ProcessPullRequest(payload []byte) string {
	var eventPayload domain.PullRequestEventPayload

	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		return "ERROR"
	}

	if eventPayload.Action == "closed" {
		base := eventPayload.PullRequest.Base.Ref
		branch := eventPayload.PullRequest.Head.Ref
		user := eventPayload.PullRequest.User.Login
		pRID := eventPayload.PullRequest.ID

		log.Printf("Pull Request Recibido:\nID:%d\nBase:%s\nHead:%s\nUser:%s", pRID, base, branch, user)
	} else {
		log.Printf("Pull Request Action no es Closed: %s", eventPayload.Action)
	}

	base := eventPayload.PullRequest.Base.Ref
	head := eventPayload.PullRequest.Head.Ref
	html_url := eventPayload.PullRequest.HTML_URL
	user := eventPayload.PullRequest.User.Login
	repository_full_name := eventPayload.Repository.FullName

	return generateDiscordMessage(base, head, html_url, user, repository_full_name)
}
