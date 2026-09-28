package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/republic/hyperdrive-calibration/internal/models"
	"github.com/segmentio/kafka-go"
)

// StartHoloNetListener connects to Kafka and begins listening for hull completion events.
func StartHoloNetListener(ctx context.Context, kafkaBroker string, topic string) {
	// Initialize the Kafka Reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kafkaBroker},
		Topic:     topic,
		Partition: 0,    // For simplicity, we read from partition 0. In prod, we'd use consumer groups.
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	// Ensure the reader closes when the service shuts down
	defer reader.Close()

	log.Printf("📡 Hyperdrive Calibration Service listening to Kafka topic: %s", topic)

	// The infinite listening loop
	for {
		// ReadMessage blocks until a new message arrives from the HoloNet
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("❌ Error reading from HoloNet: %v", err)
			continue
		}
		// Translate the raw JSON bytes into our Go struct
		var event models.HullCompletedEvent
		if err := json.Unmarshal(m.Value, &event); err != nil {
			log.Printf("❌ Failed to decode event payload: %v", err)
			continue
		}

		// We got a ship! Hand it off to the Navicomputer (to be built in Mission 2.2)
		processHyperspaceCalculation(event)
	}
}

// processHyperspaceCalculation handles the incoming event.
func processHyperspaceCalculation(event models.HullCompletedEvent) {
	log.Printf("🚀 Received Hull Completion for Ship: %s (Class: %s). Initiating Navicomputer...", event.ManifestID, event.ClassName)

	// TODO: Mission 2.2 - We will plug the heavy concurrent calculation logic in here!
	fmt.Printf("Calculating route for %s through the Unknown Regions...\n", event.ClassName)
}
