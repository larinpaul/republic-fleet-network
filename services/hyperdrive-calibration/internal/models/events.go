package models

import "time"

// HullCompletedEvent represents the JSON payload received from the Kotlin Hull Fabrication service.
// The `json:"..."` tags tell Go exactly which JSON keys map to which struct fields.
type HullCompletedEvent struct {
	ManifestID  string    `json:"manifestId"`
	ClassName   string    `json:"className"`
	HullAlloy   string    `json:"hullAlloy"`
	CompletedAt time.Time `json:"completedAt"`
}

// HyperspaceRoute represents the calculated route we will eventually return to the fleet.
type HyperspaceRoute struct {
	ManifestID   string    `json:"manifestId"`
	RouteID      string    `json:"routeId"`
	Parsecs      float64   `json:"parsecs"`
	CalculatedAt time.Time `json:"calculatedAt"`
}
