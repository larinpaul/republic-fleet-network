package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/republic/hyperdrive-calibration/internal/calculator"
	"github.com/republic/hyperdrive-calibration/internal/consumer"
	"github.com/republic/hyperdrive-calibration/internal/producer"
)

func main() {
	log.Println("🌌 Igniting Hyperdrive Calibration Service...")

	kafkaBroker := getEnv("KAFKA_BROKER", "localhost:9092")
	incomingTopic := "republic.ship.hull.completed"
	outgoingTopic := "republic.ship.route.calculated" // The new broadcast channel!
	httpPort := getEnv("HTTP_PORT", "8082")

	// --- 1. Initialize the HoloNet Broadcaster (Kafka Producer) ---
	broadcaster := producer.NewHoloNetBroadcaster(kafkaBroker, outgoingTopic)
	defer broadcaster.Close() // Ensure it closes gracefully on shutdown

	// --- 2. Initialize and Start the Navicomputer ---
	navi := calculator.NewNavicomputer(5, broadcaster)
	navi.Start()

	// --- 3. Start the HoloNet Listener (Kafka Consumer) ---
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go consumer.StartHoloNetListener(ctx, kafkaBroker, incomingTopic, navi)

	// --- 4. Start the HTTP API Server ---
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status": "ONLINE", "service": "Hyperdrive Calibration", "navicomputer_workers": 5}`)
	})

	server := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

	go func() {
		log.Printf("🌐 HTTP API listening on port %s", httpPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ HTTP server failed: %v", err)
		}
	}()

	// --- 5. Graceful Shutdown ---
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down the fleet gracefully...")
	cancel()
	server.Shutdown(context.Background())
	log.Println("May the Force be with you. Goodbye.")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
