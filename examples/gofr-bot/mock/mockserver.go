package main

import (
	"errors"
	"math/rand"

	"gofr.dev/pkg/gofr"
)

// SocialMediaPost represents the structure of a social media post
type SocialMediaPost struct {
	Content string `json:"content"`
}

// MockPublishHandler handles publishing posts to social media (mocked)
func MockPublishHandler(c *gofr.Context) (interface{}, error) {
	// Initialize a SocialMediaPost struct
	var post SocialMediaPost

	// Bind the incoming request payload to the struct
	if err := c.Bind(&post); err != nil {
		c.Errorf("Failed to bind request: %v", err)
		return nil, errors.New("invalid request payload")
	}

	// Log the mock publishing action
	c.Infof("Publishing to %s:", "Twitter")

	// Return a success response
	response := map[string]string{
		"status":  "success",
		"message": "Post successfully published to Twitter!",
	}
	return response, nil
}

// MockTrendingTopics contains a list of mock trending topics
var MockTrendingTopics = []string{
	"Latest in Go",
	"GoFr Framework Release",
	"GoFr vs Fiber",
	"Golang HTTP service Simplified",
	"Microservices Best Practices",
	"Go datasource",
	"Go Performance Tuning",
	"Serverless Architecture Trends",
}

// TrendingHandler returns a random trending topic from the mock list
func TrendingHandler(c *gofr.Context) (interface{}, error) {
	// Select a random trending topic
	randomIndex := rand.Intn(len(MockTrendingTopics))
	topic := MockTrendingTopics[randomIndex]

	// Return the selected topic as a JSON response
	return topic, nil
}

func main() {
	// Create a new GoFr application
	app := gofr.New()

	// Add a route for the mock publish endpoint
	app.POST("/api/twitter/publish", MockPublishHandler)

	// Add a route for the mock trending topics endpoint
	app.GET("/api/twitter/trending", TrendingHandler)

	// Run the application
	app.Run()
}
