package calculator

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/republic/hyperdrive-calibration/internal/models"
	"github.com/republic/hyperdrive-calibration/internal/producer"
)

type Navicomputer struct {
	inputChan   chan models.HullCompletedEvent
	workerCount int
	broadcaster *producer.HoloNetBroadcaster // <-- Injected the Broadcaster
}

// Updated constructor to accept the broadcaster
func NewNavicomputer(workerCount int, broadcaster *producer.HoloNetBroadcaster) *Navicomputer {
	return &Navicomputer{
		inputChan:   make(chan models.HullCompletedEvent, 100),
		workerCount: workerCount,
		broadcaster: broadcaster,
	}
}

func (n *Navicomputer) Start() {
	log.Printf("🧠 Igniting Navicomputer with %d concurrent astromech workers...", n.workerCount)
	for i := 1; i <= n.workerCount; i++ {
		go n.worker(i)
	}
}

func (n *Navicomputer) SubmitRoute(event models.HullCompletedEvent) {
	n.inputChan <- event
}

func (n *Navicomputer) worker(id int) {
	for event := range n.inputChan {
		log.Printf("⚙️ Worker %d calculating route for ship %s (%s)...", id, event.ManifestID, event.ClassName)

		// Simulate heavy CPU computation
		time.Sleep(2 * time.Second)
		parsecs := 8.0 + rand.Float64()*4.0

		route := models.HyperspaceRoute{
			ManifestID:   event.ManifestID,
			RouteID:      fmt.Sprintf("RTE-%d-%d", time.Now().Unix(), id),
			Parsecs:      parsecs,
			CalculatedAt: time.Now(),
		}

		log.Printf("✅ Worker %d FINISHED! Ship %s route calculated: %.2f parsecs.", id, event.ClassName, route.Parsecs)

		// 🚀 BROADCAST TO THE HOLONET!
		if err := n.broadcaster.PublishRouteCalculated(context.Background(), route); err != nil {
			log.Printf("⚠️ Worker %d failed to broadcast route for ship %s", id, event.ManifestID)
		}
	}
}
