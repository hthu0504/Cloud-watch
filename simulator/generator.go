package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateEvent(server Server) TelemetryEvent {
	return TelemetryEvent{
		EventID:   fmt.Sprintf("evt-%d", rand.Int63()),
		ServerID:  server.ID,
		Timestamp: time.Now(),
		CPU:       30 + rand.Float64()*40,
		Memory:    40 + rand.Float64()*35,
		Latency:   10 + rand.Float64()*40,
		Status:    "healthy",
	}
}