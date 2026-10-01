package calculator

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/republic/hyperdrive-calibration/internal/models"
)

// Navicomputer manages a pool of concurrent workers to calculate hyperspace routes.
type Navicomputer struct {
	inputChan   chan models.HullCompletedEvent // A channel is a typed, concurrency-safe pipe that lets goroutines send and receive values.
	workerCount int
}

// NewNavicomputer initializes the engine with a buffered channel.
// The buffer (100) ensures the Kafka consumer never blocks, even if all workers are busy.
func NewNavicomputer(workerCount int) *Navicomputer {
	return &Navicomputer{
		inputChan:   make(chan models.HullCompletedEvent, 100),
		workerCount: workerCount,
	}
}

// Start ignites the worker pool.
func (n *Navicomputer) Start() {
	log.Printf("🧠 Igniting Navicomputer with %d concurrent astromech workers...", n.workerCount)
	for i := 1; i <= n.workerCount; i++ {
		go n.worker(i)
	}
}

// SubmitRoute hands a new ship event to the worker pool.
func (n *Navicomputer) SubmitRoute(event models.HullCompletedEvent) {
	// This is non-blocking as long as the channel buffer isn't full
	n.inputChan <- event
}

// worker is the individual Goroutine that does the heavy mathematical lifting.
func (n *Navicomputer) worker(id int) {
	for event := range n.inputChan {
		log.Printf("⚙️ Worker %d calculating route for ship %s (%s)...", id, event.ManifestID, event.ClassName)

		// --- SIMULATED HEAVY CPU COMPUTATION ---
		// In reality, this would be complex spatial math, checking star charts,
		// and avoiding Imperial blockades. We simulate it with a sleep and random math.
		time.Sleep(2 * time.Second) // Simulate 2 seconds of intense calculation

		// Calculate the parsecs (Make it under 12 parsecs, obviously)
		parsecs := 8.0 + rand.Float64()*4.0

		// Generate the final route object
		route := models.HyperspaceRoute{
			ManifestID:   event.ManifestID,
			RouteID:      fmt.Sprintf("RTE-%d-%d", time.Now().Unix(), id),
			Parsecs:      parsecs,
			CalculatedAt: time.Now(),
		}

		log.Printf("✅ Worker %d FINISHED! Ship %s route calculated: %.2f parsecs. Route ID: %s",
			id, event.ClassName, route.Parsecs, route.RouteID)

		// TODO (Mission 2.3): We will eventually publish this 'route.calculated' event
		// back to Kafka so the rest of the fleet knows the ship is ready to jump!
	}
}
