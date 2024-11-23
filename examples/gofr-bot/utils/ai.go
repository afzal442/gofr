package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gofr.dev/pkg/gofr"
)

// Simulate AI-based post generation
/* func GeneratePost(trendingTopics []string, updates string) string {
	post := fmt.Sprintf("Discover the latest in GoFr! 🚀\n%s\n#GoLang #Microservices #TrendingTopics", updates)
	if len(trendingTopics) > 0 {
		post += "\nTrending Topics: " + strings.Join(trendingTopics, ", ")
	}
	return post
} */

/* func GeneratePost(c *gofr.Context, trendingTopics, updates string) string {
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
} */

type OllamaResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Message   struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	Done          bool   `json:"done"`
	DoneReason    string `json:"done_reason,omitempty"`
	TotalDuration int64  `json:"total_duration,omitempty"`
}

func GeneratePost(c *gofr.Context, trendingTopics, updates string) string {
	// Define the API endpoint and model
	apiURL := "http://localhost:11434/api/chat"
	model := "gemma2:2b"

	// Prepare the prompt with updates and trending topics
	prompt := "Write a LinkedIn post about GoFr, a Go-based framework for microservices. Include: " +
		"\n- Recent updates: " + updates +
		"\n- Trending topics: " + trendingTopics +
		"\n- Use a friendly, professional tone with hashtags like #GoLang and #Microservices. Link to: https://github.com/gofr-dev/gofr/releases/tag/latest"

	// Create the request payload
	payload := map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	// Marshal the payload into JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		c.Errorf("Failed to marshal payload: %v", err)
		return ""
	}

	// Make the HTTP POST request to the Ollama API
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		c.Errorf("Failed to call Ollama API: %v", err)
		return ""
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Errorf("Failed to read response body: %v", err)
		return ""
	}

	// Log the raw response for debugging
	c.Infof("Raw response from Ollama API: %s", body)

	// Process NDJSON response
	scanner := bufio.NewScanner(bytes.NewReader(body))
	var fullContent strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		var responseData OllamaResponse

		// Parse each line as JSON
		if err := json.Unmarshal([]byte(line), &responseData); err != nil {
			c.Errorf("Failed to decode line: %v", err)
			continue
		}

		// Append content if available
		if strings.TrimSpace(responseData.Message.Content) != "" {
			fullContent.WriteString(responseData.Message.Content)
		}
	}

	if err := scanner.Err(); err != nil {
		c.Errorf("Failed to scan response: %v", err)
		return ""
	}

	// Return the aggregated content
	content := fullContent.String()
	if strings.TrimSpace(content) == "" {
		c.Error("Received empty content from Ollama API")
		return "The AI was unable to generate content. Please try again with a more specific prompt."
	}

	return content
}

// Simulate email content generation
func GenerateEmail(recipient string, trendingTopics, updates string) string {
	parts := strings.Split(recipient, "@")
	/* 	if len(parts) != 2 {
		return "invalid email format"
	} */
	return fmt.Sprintf("Hi %s,\n\nExplore the power of GoFr for microservices!\n%s\n\nBest,\nThe GoFr Team", parts[0], updates)
}
