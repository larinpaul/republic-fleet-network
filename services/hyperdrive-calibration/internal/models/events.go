package models

import "time"

// HullCompletedEvent (Incoming from Kotlin)
type HullCompletedEvent struct {
	ManifestID  string    `json:"manifestId"`
	ClassName   string    `json:"className"`
	HullAlloy   string    `json:"hullAlloy"`
	CompletedAt time.Time `json:"completedAt"`
}

// HyperspaceRoute (Internal calculation result)
type HyperspaceRoute struct {
	ManifestID   string    `json:"manifestId"`
	RouteID      string    `json:"routeId"`
	Parsecs      float64   `json:"parsecs"`
	CalculatedAt time.Time `json:"calculatedAt"`
}

// RouteCalculatedEvent (Outgoing to Kafka - The Broadcast)
// This is what the rest of the fleet (like the Astromech Provisioning service) will listen for.
type RouteCalculatedEvent struct {
	ManifestID   string    `json:"manifestId"`
	RouteID      string    `json:"routeId"`
	Parsecs      float64   `json:"parsecs"`
	CalculatedAt time.Time `json:"calculatedAt"`
}
