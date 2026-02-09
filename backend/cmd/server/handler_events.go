package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "there is a problem with SSE", http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	for {
		select {
		case msg := <-cfg.sseChan:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()
		case <-ctx.Done():
			return
		}
	}
}
