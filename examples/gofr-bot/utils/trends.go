package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"gofr.dev/pkg/gofr"
)

// FetchTrendingTopics fetches a random trending topic from the mock server

// MockTrendingTopics is the list of mock topics we return from our mock server.
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

// TrendingHandler will return a random topic from our mock list of topics.
func TrendingHandler(c *gofr.Context) (interface{}, error) {
	// Select a random trending topic
	randomIndex := rand.Intn(len(MockTrendingTopics))
	topic := MockTrendingTopics[randomIndex]

	// Return the selected topic as a JSON response
	return topic, nil
}

func FetchTrendingTopics(c *gofr.Context) string {
	// Send a request to the mock server to fetch trending topics
	resp, err := http.Get("http://localhost:5000/api/twitter/trending")
	if err != nil || resp.StatusCode != http.StatusOK {
		c.Errorf("Failed to fetch trending topics: %v", err)
		return "GoFr Latest Update"
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Errorf("Failed to read response body: %v", err)
		return "GoFr Latest Update"
	}

	// Define a map to hold the response
	var result map[string]interface{}

	// Decode the response body into the map
	if err := json.Unmarshal(body, &result); err != nil {
		c.Errorf("Failed to decode trending topics: %v", err)
		return "GoFr Latest Update"
	}

	// Extract the 'data' field from the map and assert it as a string
	topic, ok := result["data"].(string)
	if !ok {
		c.Errorf("Failed to extract 'data' from the response")
		return "GoFr Latest Update"
	}

	return topic
}

type GitHubRelease struct {
	TagName     string `json:"tag_name"`
	Name        string `json:"name"`
	Body        string `json:"body"`
	PublishedAt string `json:"published_at"`
}

func FetchgofrUpdates(c *gofr.Context) (string, error) {
	// Define the GitHub API URL for GoFr releases
	apiURL := "https://api.github.com/repos/gofr-dev/gofr/releases"

	// Create a new HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch releases: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 HTTP status
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned status: %s", resp.Status)
	}

	// Parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var releases []GitHubRelease
	err = json.Unmarshal(body, &releases)
	if err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Ensure there are releases available
	if len(releases) == 0 {
		return "No releases found for GoFr.", nil
	}

	// Fetch details of the gofr latest release
	gofr := releases[0]
	gofrContent := fmt.Sprintf(
		"gofr Content: %s\nTag: %s\nDetails:\n%s",
		gofr.Name, gofr.TagName, gofr.Body,
	)

	return gofrContent, nil
}
