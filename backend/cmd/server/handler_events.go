package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tfriezzz/tourtap/internal/auth"
)

func (cfg *apiConfig) handlerEvents(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")

	cookie, err := r.Cookie("ssh_auth")
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "unauthorized", err)
		return
	}
	if _, err := auth.ValidateJWT(cookie.Value, cfg.jwtSecret); err != nil {
		respondWithError(w, http.StatusUnauthorized, "unauthorized", err)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	// w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "there is a problem with SSE", http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	fmt.Fprintf(w, ": connected\n\n")
	flusher.Flush()

	heartBeat := time.NewTicker(25 * time.Second)
	defer heartBeat.Stop()

	for {
		select {
		case msg := <-cfg.sseChan:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()

		case <-heartBeat.C:
			fmt.Fprintf(w, ": ping\n\n")
			flusher.Flush()

		case <-ctx.Done():
			// fmt.Println("SSE client disconnected")
			return
		}
	}
}
