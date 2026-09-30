package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/republic/hyperdrive-calibration/internal/consumer"
)

func main() {
	log.Println("🌌 Igniting Hyperdrive Calibration Service...")

	// --- 1. Configuration ---
	// In a real app, we'd read these from environment variables or a config file
	kafkaBroker := getEnv("KAFKA_BROKER", "localhost:9092")
	kafkaTopic := "republic.ship.hull.completed"
	httpPort := getEnv("HTTP_PORT", "8082")

	// --- 2. Start the HoloNet Listener (Kafka Consumer) ---
	// We run this in a Goroutine so it doesn't block the HTTP server from starting
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go consumer.StartHoloNetListener(ctx, kafkaBroker, kafkaTopic)

	// --- 3. Start the HTTP API Server ---
	// This is the REST API where pilots can query the status of their route calculations
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status": "ONLINE", "service": "Hyperdrive Calibration", "message": "Navicomputer spooled and ready."}`)
	})

	server := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

	// Start the HTTP server in a Goroutine
	go func() {
		log.Printf("🌐 HTTP API listening on port %s", httpPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ HTTP server failed: %v", err)
		}
	}()

	// --- 4. Graceful Shutdown ---
	// Wait for an interrupt signal (Ctrl+C) to shut down both the HTTP server and Kafka listener cleanly
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down the fleet gracefully...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	cancel()                     // Stops the Kafka listener
	server.Shutdown(shutdownCtx) // Stops the HTTP server

	log.Println("May the Force be with you. Goodbye.")
}

// getEnv is a helper to read environment variables with a fallback default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
