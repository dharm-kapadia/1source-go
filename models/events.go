// Models package contains the models for the application
package models

type (
	Events []event

	event struct {
		EventId       uint64 `json:"eventId"`
		EventType     string `json:"eventType"`
		EventDateTime string `json:"eventDateTime"`
		ResourceUri   string `json:"resourceUri"`
	}
)
