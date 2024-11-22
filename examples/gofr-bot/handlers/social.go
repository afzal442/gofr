package handlers

import (
	"gofr.dev/pkg/gofr"
	"gofr-bot/utils"
)

// GeneratePostHandler drafts a social media post
func GeneratePostHandler(c *gofr.Context) (interface{}, error) {
	trendingTopics := utils.FetchTrendingTopics() // Get trending topics
	latestUpdates := utils.FetchLatestUpdates()  // Pull GoFr updates

	post := utils.GeneratePost(trendingTopics, latestUpdates) // AI-generated post

	// Save post draft to Redis for review
	utils.SaveDraftToRedis(c, "social_post", post)
	return map[string]string{"draft": post}, nil
}

// ApprovePostHandler publishes an approved post
func ApprovePostHandler(c *gofr.Context) (interface{}, error) {
	post, err := utils.GetDraftFromRedis(c, "social_post")
	if err != nil {
		return nil, err
	}

	err = utils.PublishToSocialMedia(post) // Publish to LinkedIn/Twitter
	if err != nil {
		return nil, err
	}

	return map[string]string{"status": "Post published successfully"}, nil
}
