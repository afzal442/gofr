package utils

import (
	"log"

	"gofr.dev/pkg/gofr"
)

func PublishToSocialMedia(post string) error {
	// Placeholder: Use LinkedIn and Twitter APIs
	log.Println("Publishing post: ", post)
	return nil
}

func SaveDraftToRedis(c *gofr.Context, key string, draft string) {
	c.Redis.Set(c, key, draft, 0)
}

func GetDraftFromRedis(c *gofr.Context, key string) (string, error) {
	return c.Redis.Get(c, key).Result()
}
