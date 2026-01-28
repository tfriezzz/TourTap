package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tfriezzz/tourtap/internal/database"
	"github.com/tfriezzz/tourtap/internal/pubsub"
)

type apiConfig struct {
	db   *database.Queries
	port string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("cannot load .env file")
	}
	port := os.Getenv("TOURTAP_PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}
	dbURL := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	if err != nil {
		log.Printf("can't connect to database: %v\n", err)
	}

	apiCfg := apiConfig{
		port: port,
		db:   dbQueries,
	}

	if err := pubsub.Init(); err != nil {
		fmt.Println("nope")
		log.Fatal(err)
	}
	defer pubsub.Close()

	pubsub.Start(func(event pubsub.Event) {
		switch event.Type {
		case "group_created":
			payload := Group{}
			if err := json.Unmarshal(event.Data, &payload); err != nil {
				log.Printf("could not unmarshal payload: %v", err)
			}
			log.Printf("New group request: %v pax for tour %v on %v", payload.Pax, payload.RequestedTourID, payload.RequestedDate)

		case "group_accepted":
			payload := Group{}
			if err := json.Unmarshal(event.Data, &payload); err != nil {
				log.Printf("could not unmarshal payload")
			}
			log.Printf("Group %v for tour %v on %v accepted", payload.Email, payload.RequestedTourID, payload.RequestedDate)
			log.Println("*sending payment information*")

		case "group_declined":
			// TODO:
		}
	})

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
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
