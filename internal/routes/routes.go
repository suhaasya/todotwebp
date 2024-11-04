package routes

import (
	"net/http"
	"todotwebp/internal/handlers"
)

func InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/images", handlers.ImageHandler)
	return mux
}
