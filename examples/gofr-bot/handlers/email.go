package handlers

import (
	"gofr-bot/utils"

	"gofr.dev/pkg/gofr"
)

// EmailOutreachHandler drafts and sends outreach emails
func EmailOutreachHandler(c *gofr.Context) (interface{}, error) {
	recipients := utils.FetchTargetRecipients()   // Fetch target emails
	trendingTopics := utils.FetchTrendingTopics() // Get trending topics
	latestUpdates := utils.FetchLatestUpdates()   // GoFr updates

	for _, recipient := range recipients {
		email := utils.GenerateEmail(recipient, trendingTopics, latestUpdates)
		err := utils.SendEmail(email)
		if err != nil {
			return nil, err
		}
	}

	return map[string]string{"status": "Emails sent successfully"}, nil
}
