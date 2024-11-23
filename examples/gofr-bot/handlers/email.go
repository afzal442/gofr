package handlers

import (
	"fmt"
	"gofr-bot/utils"

	"gofr.dev/pkg/gofr"
)

// EmailOutreachHandler drafts and sends outreach emails
func EmailOutreachHandler(c *gofr.Context) (interface{}, error) {
	recipients := utils.FetchTargetRecipients(c)   // Fetch target emails
	trendingTopics := utils.FetchTrendingTopics(c) // Get trending topics
	gofrUpdates, err := utils.FetchgofrUpdates(c)  // GoFr updates

	if err != nil {
		return "", fmt.Errorf("failed to fetch updates: %w", err)
	}

	for _, recipient := range recipients {
		email := utils.GenerateEmail(recipient, trendingTopics, gofrUpdates)
		err := utils.SendEmail(c, trendingTopics, recipient, email)
		if err != nil {
			return nil, err
		}
	}

	return map[string]string{"status": "Emails sent successfully"}, nil
}
