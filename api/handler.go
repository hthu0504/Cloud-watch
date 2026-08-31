package main

import (
	"encoding/json"
	"net/http"
)

type TelemetryEvent struct {
	ServerID string  `json:"server_id"`
	CPU      float64 `json:"cpu"`
	Memory   float64 `json:"memory"`
	Latency  float64 `json:"latency"`
	Status   string  `json:"status"`
}

func telemetryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var event TelemetryEvent

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	err = publishEvent(event)
	if err != nil {
		http.Error(w, "failed to publish event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "event published",
	})
}