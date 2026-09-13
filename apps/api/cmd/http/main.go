package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type healthResponse struct {
	Status string `json:"status"`
}

// @title			Conductor API
// @version			0.1.0
// @description 	Distributed workflow engine API.
// @BasePath 		/
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(healthResponse{Status: "ok"}); err != nil {
			log.Printf("failed to encode health response: %v", err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
