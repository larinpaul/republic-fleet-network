package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/republic/hyperdrive-calibration/internal/calculator"
	"github.com/republic/hyperdrive-calibration/internal/models"
	"github.com/segmentio/kafka-go"
)

// StartHoloNetListener connects to Kafka and feeds events directly to the Navicomputer.
func StartHoloNetListener(ctx context.Context, kafkaBroker string, topic string, navicomputer *calculator.Navicomp) {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kafkaBroker},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	defer reader.Close()

	log.Printf("📡 Hyperdrive Calibration listening to Kafka topic: %s", topic)

	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("❌ Error reading from HoloNet: %v", err)
			continue
		}

		var event models.HullCompletedEvent
		if err := json.Unmarshal(m.Value, &event); err != nil {
			log.Printf("❌ Failed to decode event payload: %v", err)
			continue
		}

		// HAND OFF TO THE NAVICOMPUTER!
		// The consumer doesn't calculate; it just routes the traffic.
		navicomputer.SubmitRoute(event)
	}
}
