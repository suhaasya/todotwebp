package main

import (
	"fmt"
	"log"
	"net/http"
	"todotwebp/internal/routes"

	"github.com/rs/cors"
)

func main() {
	// Get the handler with routes
	handler := routes.InitRoutes()

	// Set up CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Replace "*" with specific origins for security
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Wrap handler with CORS and start the server
	fmt.Println("Running Server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(handler)))
}
