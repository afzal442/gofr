package utils

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
)

// Simulate AI-based post generation
/* func GeneratePost(trendingTopics []string, updates string) string {
	post := fmt.Sprintf("Discover the latest in GoFr! 🚀\n%s\n#GoLang #Microservices #TrendingTopics", updates)
	if len(trendingTopics) > 0 {
		post += "\nTrending Topics: " + strings.Join(trendingTopics, ", ")
	}
	return post
} */

func GeneratePost(trendingTopics []string, updates string) string {
	// Initialize GEMINI AI client
	ctx := context.Background()
	client, err := genai.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create GEMINI client: %v", err)
	}

	// Prepare prompt with updates and trending topics
	prompt := "Write a LinkedIn post about GoFr, a Go-based framework for microservices. Include: " +
		"\n- Recent updates: " + updates +
		"\n- Trending topics: " + strings.Join(trendingTopics, ", ") +
		"\n- Use a friendly, professional tone with hashtags like #GoLang and #Microservices."

	// Use GEMINI AI to generate the content
	model := client.GenerativeModel("gemini-1.5-flash")
	iter := model.GenerateContentStream(ctx, genai.Text(prompt))

	var postBuilder strings.Builder
	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error generating content: %v", err)
		}
		postBuilder.WriteString(resp.PromptFeedback.BlockReason.String())
	}

	// Return the generated post
	return postBuilder.String()
}

// Simulate email content generation
func GenerateEmail(recipient string, trendingTopics []string, updates string) string {
	return fmt.Sprintf("Hi %s,\n\nExplore the power of GoFr for microservices!\n%s\n\nBest,\nThe GoFr Team", recipient, updates)
}
