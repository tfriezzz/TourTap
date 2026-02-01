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
	db        *database.Queries
	jwtSecret string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("cannot load .env file")
	}
	const jwtSecret = os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}
	// TODO: add jwtSecret to .env
	const port = os.Getenv("TOURTAP_PORT")
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
		db:        dbQueries,
		jwtSecret: jwtSecret,
	}

	// if err := templates.Load(); err != nil {
	// 	log.Fatal(err)
	// }

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
			log.Printf("New group request: %v pax for tour %v on %v pending", payload.Pax, payload.RequestedTourID, payload.RequestedDate)

		case "group_accepted":
			payload := Group{}
			if err := json.Unmarshal(event.Data, &payload); err != nil {
				log.Printf("could not unmarshal payload")
			}
			log.Printf("Group %v for tour %v on %v accepted", payload.Email, payload.RequestedTourID, payload.RequestedDate)
			log.Println("*sending payment information*")

		case "group_declined":
			payload := Group{}
			if err := json.Unmarshal(event.Data, &payload); err != nil {
				log.Printf("could not unmarshal payload")
			}
			log.Printf("Group %v for tour %v on %v declined", payload.Email, payload.RequestedTourID, payload.RequestedDate)
			log.Println("*Sending decline mail*")

		case "group_confirmed":
			payload := Group{}
			if err := json.Unmarshal(event.Data, &payload); err != nil {
				log.Printf("could not unmarshal payload")
			}
			log.Printf("Group %v for tour %v on %v confirmed", payload.Email, payload.RequestedTourID, payload.RequestedDate)
			log.Printf("*sending reciept*")

		}
	})

	mux := http.NewServeMux()

	mux.Handle("GET /api/health", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerReadiness)))
	mux.Handle("POST /api/groups/create", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerGroupsCreate)))
	mux.Handle("POST /admin/tours/create", http.StripPrefix("/admin/", http.HandlerFunc(apiCfg.handlerToursCreate)))
	mux.Handle("GET /api/bookings", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerBookingsGet)))
	mux.Handle("POST /admin/reset-groups", http.StripPrefix("/admin/", http.HandlerFunc(apiCfg.handlerGroupsReset)))
	mux.Handle("GET /api/tours", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerToursGet)))
	mux.Handle("GET /api/pending", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerPending)))
	mux.Handle("PUT /api/groups/{groupID}/accept", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerGroupsAccept)))
	mux.Handle("PUT /api/groups/{groupID}/decline", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerGroupsDecline)))
	mux.Handle("POST /webhooks/payment", http.HandlerFunc(apiCfg.handlerPaymentWebhook))

	mux.Handle("POST /api/users", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerUsersCreate)))
	mux.Handle("PUT /api/users", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerUpdateCredentials)))
	mux.Handle("POST /api/login", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerLogin)))
	mux.Handle("POST /api/refresh", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerRefresh)))
	mux.Handle("POST /api/revoke", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerRevoke)))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("serving on: http://localhost:%s/\n", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
