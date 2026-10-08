package producer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/republic/hyperdrive-calibration/internal/models"
	"github.com/segmentio/kafka-go"
)

// HoloNetBroadcaster handles publishing events back to the Republic HoloNet.
type HoloNetBroadcaster struct {
	writer *kafka.Writer
}

// NewHoloNetBroadcaster initializes the Kafka Writer.
func NewHoloNetBroadcaster(broker string, topic string) *HoloNetBroadcaster {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{}, // Distributes load evenly across partitions
	}
	return &HoloNetBroadcaster{writer: writer}
}

// PublishRouteCalculated broadcasts the calculated hyperspace route to the fleet.
func (b *HoloNetBroadcaster) PublishRouteCalculated(ctx context.Context, route models.HyperspaceRoute) error {
	// 1. Marshal the Go struct into a JSON byte array
	payload, err := json.Marshal(route)
	if err != nil {
		log.Printf("❌ Failed to marshal route event: %v", err)
		return err
	}

	// 2. Construct the Kafka Message
	msg := kafka.Message{
		Key:   []byte(route.ManifestID), // Keep events for the same ship in the same partition
		Value: payload,
	}

	// 3. Fire the laser! Write the message to Kafka
	err = b.writer.WriteMessages(ctx, msg)
	if err != nil {
		log.Printf("❌ Failed to broadcast route to HoloNet: %v", err)
		return err
	}

	log.Printf("📡 Broadcasted Route ID %s for Ship %s to HoloNet!", route.RouteID, route.ManifestID)
	return nil
}

// Close gracefully shuts down the Kafka writer when the service stops.
func (b *HoloNetBroadcaster) Close() error {
	return b.writer.Close()
}
