package utils

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"gofr.dev/pkg/gofr"
	"google.golang.org/api/option"
)

// Simulate AI-based post generation
/* func GeneratePost(trendingTopics []string, updates string) string {
	post := fmt.Sprintf("Discover the latest in GoFr! 🚀\n%s\n#GoLang #Microservices #TrendingTopics", updates)
	if len(trendingTopics) > 0 {
		post += "\nTrending Topics: " + strings.Join(trendingTopics, ", ")
	}
	return post
} */

func GeneratePost(c *gofr.Context, trendingTopics, updates string) string {
	// Initialize GEMINI AI client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(""))
	if err != nil {
		c.Errorf("Failed to create GEMINI client: %v", err)
	}

	// Prepare prompt with updates and trending topics
	prompt := "Write a LinkedIn post about GoFr, a Go-based framework for microservices. Include: " +
		"\n- Recent updates: " + updates +
		"\n- Trending topics: " + trendingTopics +
		"\n- Use a friendly, professional tone with hashtags like #GoLang and #Microservices. and link to https://github.com/gofr-dev/gofr/releases/tag/ + realeas version"

	// Use GEMINI AI to generate the content
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		c.Error(err)
	}

	var postBuilder strings.Builder

	if resp == nil {
		c.Error("Empty response from GEMINI AI")
		return ""
	}

	// Iterate through candidates in the response
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				// Convert part to string dynamically if it doesn't have a direct string field
				partStr := fmt.Sprintf("%v", part) // Replace this with the correct field/method if available
				postBuilder.WriteString(partStr)
			}
		}
	}

	// Return the generated post
	return postBuilder.String()
}

// Simulate email content generation
func GenerateEmail(recipient string, trendingTopics, updates string) string {
	parts := strings.Split(recipient, "@")
	/* 	if len(parts) != 2 {
		return "invalid email format"
	} */
	return fmt.Sprintf("Hi %s,\n\nExplore the power of GoFr for microservices!\n%s\n\nBest,\nThe GoFr Team", parts[0], updates)
}
