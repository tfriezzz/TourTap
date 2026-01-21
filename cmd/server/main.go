package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	// db   *database.Queries
	port string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("cannot load .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	cfg := apiConfig{
		port: port,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", http.HandlerFunc(cfg.handlerReadiness))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("serving on: http://localhost:%s/\n", port)
	log.Fatal(srv.ListenAndServe())
}
