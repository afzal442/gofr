package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"gofr.dev/pkg/gofr"
)

type SocialMediaPost struct {
	Content string `json:"content"`
}

func PublishToSocialMedia(c *gofr.Context, content string) error {
	// Prepare the request payload
	post := SocialMediaPost{
		Content: content,
	}

	postData, err := json.Marshal(post)
	if err != nil {
		return fmt.Errorf("failed to serialize post: %w", err)
	}

	// Send the request to the mock server
	resp, err := http.Post("http://localhost:5000/api/twitter/publish", "application/json", bytes.NewBuffer(postData))
	if err != nil {
		return fmt.Errorf("failed to publish post: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("server returned status: %s", resp.Status)
	}

	return nil
}

func SaveDraftToRedis(c *gofr.Context, key string, draft string) {
	c.Redis.Set(c, key, draft, 0)
}

func GetDraftFromRedis(c *gofr.Context, key string) (string, error) {
	return c.Redis.Get(c, key).Result()
}
