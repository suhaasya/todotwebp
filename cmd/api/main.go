package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todotwebp/internal/routes"

	"github.com/joho/godotenv"
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

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Wrap handler with CORS and start the server
	fmt.Printf("Running Server on port :%v \n", port)
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(handler)))
}
