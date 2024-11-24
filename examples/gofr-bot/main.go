package main

import (
	"gofr-bot/handlers"
	"net/http"

	gofrHTTP "gofr.dev/pkg/gofr/http"

	"gofr.dev/pkg/gofr"
)

func corsMiddleware() gofrHTTP.Middleware {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Add CORS headers to the response
			w.Header().Set("Access-Control-Allow-Origin", "*")                                // Allow requests from any origin. Replace '*' with a specific domain if needed.
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Allowed HTTP methods
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")     // Allowed headers

			// Proceed to the next middleware or handler
			inner.ServeHTTP(w, r)
		})
	}
}

func main() {
	// Create a new GoFr application
	app := gofr.New()

	// Define routes
	app.GET("/api/social/posts", handlers.GeneratePostHandler)
	app.POST("/api/social/approve", handlers.ApprovePostHandler)
	app.POST("/api/email/outreach", handlers.EmailOutreachHandler)

	app.UseMiddleware(corsMiddleware())

	// Run the application
	app.Run()
}
