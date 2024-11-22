package utils

import (
	"fmt"
	"strings"
)

// Simulate AI-based post generation
func GeneratePost(trendingTopics []string, updates string) string {
	post := fmt.Sprintf("Discover the latest in GoFr! 🚀\n%s\n#GoLang #Microservices #TrendingTopics", updates)
	if len(trendingTopics) > 0 {
		post += "\nTrending Topics: " + strings.Join(trendingTopics, ", ")
	}
	return post
}

// Simulate email content generation
func GenerateEmail(recipient string, trendingTopics []string, updates string) string {
	return fmt.Sprintf("Hi %s,\n\nExplore the power of GoFr for microservices!\n%s\n\nBest,\nThe GoFr Team", recipient, updates)
}
