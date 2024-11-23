package main

import (
	"gofr.dev/pkg/gofr"
	"gofr-bot/handlers"
)

func main() {
	// Create a new GoFr application
	app := gofr.New()

	// Define routes
	app.GET("/api/social/posts", handlers.GeneratePostHandler)
	app.POST("/api/social/approve", handlers.ApprovePostHandler)
	app.POST("/api/email/outreach", handlers.EmailOutreachHandler)

	// Run the application
	app.Run()
}
