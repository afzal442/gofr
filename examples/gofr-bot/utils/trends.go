package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gofr.dev/pkg/gofr"
)

// FetchTrendingTopics fetches a random trending topic from the mock server
func FetchTrendingTopics(c *gofr.Context) string {
	// Send a request to the mock server to fetch trending topics
	resp, err := http.Get("http://localhost:5000/api/twitter/trending")
	if err != nil || resp.StatusCode != http.StatusOK {
		c.Errorf("Failed to fetch trending topics: %v", err)
		return "GoFr Latest Update"
	}
	defer resp.Body.Close()

	// Decode the response into a slice of strings
	var topics string
	if err := c.Bind(&topics); err != nil {
		c.Errorf("Failed to decode trending topics: %v", err)
		return "GoFr Latest Update"
	}

	return topics
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
