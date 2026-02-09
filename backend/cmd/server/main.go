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
	sseChan   chan string
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("cannot load .env file")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}
	// TODO: add jwtSecret to .env
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

	sseChan := make(chan string, 10)

	apiCfg := apiConfig{
		db:        dbQueries,
		jwtSecret: jwtSecret,
		sseChan:   sseChan,
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
			log.Printf("New group request: %v pax for tour %v on %v pending", payload.Pax, payload.RequestedTourID, payload.RequestedDate)
			sseChan <- fmt.Sprintf("New booking: %v, for tour: %v", payload.Email, payload.RequestedTourID)

		case "group_accepted":
			payload := Group{}
			if err := json.Unmarshal(event.Data, &payload); err != nil {
				log.Printf("could not unmarshal payload")
			}
			log.Printf("Group %v for tour %v on %v accepted", payload.Email, payload.RequestedTourID, payload.RequestedDate)
			sseChan <- fmt.Sprintf("Group %v: accepted", payload.Email)
			log.Println("*sending payment information*")

		case "group_declined":
			payload := Group{}
			if err := json.Unmarshal(event.Data, &payload); err != nil {
				log.Printf("could not unmarshal payload")
			}
			log.Printf("Group %v for tour %v on %v declined", payload.Email, payload.RequestedTourID, payload.RequestedDate)
			sseChan <- fmt.Sprintf("Group %v: declined", payload.Email)
			log.Println("*Sending decline mail*")

		case "group_confirmed":
			payload := Group{}
			if err := json.Unmarshal(event.Data, &payload); err != nil {
				log.Printf("could not unmarshal payload")
			}
			// log.Printf("Group %v for tour %v on %v confirmed", payload.Email, payload.RequestedTourID, payload.RequestedDate)
			sseChan <- fmt.Sprintf("Group %v for tour %v on %v: paid and confirmed", payload.Email, payload.RequestedTourID, payload.RequestedDate)
			log.Printf("*sending reciept*")

		}
	})

	mux := http.NewServeMux()

	mux.Handle("GET /api/health", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerReadiness)))

	mux.Handle("GET /api/tours", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerToursGet)))
	mux.Handle("POST /admin/tours/create", http.StripPrefix("/admin/", http.HandlerFunc(apiCfg.handlerToursCreate)))

	mux.Handle("GET /api/bookings/tour-date", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerBookingsGet)))
	mux.Handle("GET /api/bookings/all-date", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerBookingsGetAllDate)))

	// mux.Handle("GET /api/pending", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerPending)))

	mux.Handle("POST /api/groups/create", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerGroupsCreate)))
	mux.Handle("GET /api/groups/get-pending", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerGroupsGetPending)))
	mux.Handle("PUT /api/groups/{groupID}/accept", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerGroupsAccept)))
	mux.Handle("PUT /api/groups/{groupID}/decline", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerGroupsDecline)))
	mux.Handle("POST /webhooks/payment", http.HandlerFunc(apiCfg.handlerPaymentWebhook))

	mux.Handle("POST /api/users", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerUsersCreate)))
	mux.Handle("PUT /api/users", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerUsersUpdate)))
	mux.Handle("POST /api/login", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerLogin)))
	mux.Handle("POST /api/auth/refresh", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerRefresh)))
	mux.Handle("POST /api/auth/revoke", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerRevoke)))

	mux.Handle("POST /admin/reset-groups", http.StripPrefix("/admin/", http.HandlerFunc(apiCfg.handlerGroupsReset)))
	mux.Handle("POST /api/test", http.StripPrefix("/api/", http.HandlerFunc(apiCfg.handlerFrontend)))

	mux.Handle("/events", http.HandlerFunc(apiCfg.handlerEvents))

	handler := corsMiddleware(mux)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	log.Printf("serving on: http://localhost:%s/\n", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
