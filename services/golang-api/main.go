package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
		log.Fatal(err.Error())
	}
	log.Println("database connected")
	defer func() {
		log.Println("closing database pool...")
		pool.Close()
		log.Println("database pool closed")
	}()

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

	srv := &http.Server{Addr: addr, Handler: mux}

	go func() {
		log.Printf("listening on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	<-ctx.Done()

	log.Println("shutting down gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("http shutdown error: %v", err)
	}
}
