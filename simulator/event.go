package main

import "time"

type TelemetryEvent struct {
	EventID       string    `json:"event_id"`
	ServerID      string    `json:"server_id"`
	Timestamp     time.Time `json:"timestamp"`
	CPU           float64   `json:"cpu"`
	Memory        float64   `json:"memory"`
	Latency       float64   `json:"latency"`
	Status        string    `json:"status"`
}