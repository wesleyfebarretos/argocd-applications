package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/wesleyfebarretos/argocd-applications/internal/config"
	"github.com/wesleyfebarretos/argocd-applications/internal/database"
)

func writeStatus(w http.ResponseWriter, status string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": status})
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func main() {
	config.Init()

	pool, err := database.NewPostgres(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("database connected")
	defer pool.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health/livez", func(w http.ResponseWriter, r *http.Request) {
		writeStatus(w, "live")
	})
	mux.HandleFunc("GET /health/readyz", func(w http.ResponseWriter, r *http.Request) {
		if err := pool.Ping(r.Context()); err != nil {
			writeError(w, http.StatusServiceUnavailable, "database not ready")
			return
		}

		writeStatus(w, "ready")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
