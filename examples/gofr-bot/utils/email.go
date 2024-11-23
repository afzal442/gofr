package utils

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"gofr.dev/pkg/gofr"
)

// SendEmail sends an email using the Mailgun API
func SendEmail(c *gofr.Context, subject, recipient, content string) error {

	// Load Mailgun configuration from environment variables
	domain := os.Getenv("MAILGUN_DOMAIN")         // Your Mailgun domain
	privateAPIKey := os.Getenv("MAILGUN_API_KEY") // Your Mailgun private API key
	sender := os.Getenv("MAILGUN_SENDER_EMAIL")   // The sender's email address

	// Load Mailgun configuration from environment variables
	if domain == "" || privateAPIKey == "" || sender == "" {
		log.Println("Mailgun configuration is missing")
		return fmt.Errorf("mailgun configuration is missing")
	}

	// Create a new Mailgun instance
	mg := mailgun.NewMailgun(domain, privateAPIKey)

	// Create a new email message
	message := mg.NewMessage(
		sender,    // Sender
		subject,   // Subject
		content,   // Body (text or HTML)
		recipient, // Recipient
	)

	ctx, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	_, _, err := mg.Send(ctx, message)

	if err != nil {
		c.Fatal(err)
	}

	c.Log("Email successfully sent to:", recipient)
	return nil
}

// Mock email addresses to store in Redis
var emailAddresses = []string{
	"afvn.in@gmail.com",
	/* "developer1@example.com",
	"team@techcompany.com",
	"golangdev@openmail.com",
	"contact@microservices.org",
	"info@backendworld.io",
	"support@kubernetesclub.com", */
}

// FetchTargetRecipients fetches random email addresses from Redis
func FetchTargetRecipients(c *gofr.Context) []string {
	ctx := context.Background()

	// Store random email addresses in Redis if not already set
	redisKey := "email_recipients"
	if c.Redis.Exists(ctx, redisKey).Val() == 0 {
		for _, email := range emailAddresses {
			c.Redis.SAdd(ctx, redisKey, email)
		}
		c.Redis.Expire(ctx, redisKey, 24*time.Hour) // Set an expiry of 24 hours
	}

	// Fetch email addresses from Redis
	emails, err := c.Redis.SMembers(ctx, redisKey).Result()
	if err != nil {
		c.Errorf("Failed to fetch email recipients from Redis: %v", err)
		return []string{}
	}

	// Shuffle and pick a random subset of emails to return
	rand.Shuffle(len(emails), func(i, j int) { emails[i], emails[j] = emails[j], emails[i] })

	// Limit the number of recipients returned
	maxRecipients := 3
	if len(emails) < maxRecipients {
		maxRecipients = len(emails)
	}

	return emails[:maxRecipients]
}
