package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tfriezzz/tourtap/internal/database"
)

type apiConfig struct {
	db   *database.Queries
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
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	if err != nil {
		log.Printf("can't connect to database: %v\n", err)
	}

	apiCfg := apiConfig{
		port: port,
		db:   dbQueries,
	}

	mux := http.NewServeMux()

	mux.Handle("/api/health", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerReadiness)))
	mux.Handle("POST /api/groups/create", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerGroupsCreate)))
	mux.Handle("POST /admin/tours/create", http.StripPrefix("/admin/", http.HandlerFunc(apiCfg.handlerToursCreate)))
	mux.Handle("GET /api/bookings", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerBookingsGet)))
	mux.Handle("POST /admin/reset-groups", http.StripPrefix("/admin/", http.HandlerFunc(apiCfg.handlerGroupsReset)))
	mux.Handle("GET /api/tours", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerToursGet)))
	mux.Handle("GET /api/pending", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerPending)))
	mux.Handle("PUT /api/groups/{groupID}/accept", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerGroupsAccept)))
	mux.Handle("PUT /api/groups/{groupID}/decline", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerGroupsDecline)))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("serving on: http://localhost:%s/\n", port)
	log.Fatal(srv.ListenAndServe())
}
