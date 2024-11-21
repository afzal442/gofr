package main

import (
	"gofr.dev/pkg/gofr"
	"gofr-bot/handlers"
)

func main() {
	// Create a new GoFr application
	app := gofr.New()

	// Define routes
	app.GET("/social/posts", handlers.GeneratePostHandler)
	app.POST("/social/approve", handlers.ApprovePostHandler)
	app.POST("/email/outreach", handlers.EmailOutreachHandler)
	app.GET("/analytics", handlers.AnalyticsHandler)

	// Run the application
	app.Run()
}
