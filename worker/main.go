package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type TelemetryEvent struct {
	ServerID string  `json:"server_id"`
	CPU      float64 `json:"cpu"`
	Memory   float64 `json:"memory"`
	Latency  float64 `json:"latency"`
	Status   string  `json:"status"`
}

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "telemetry",
		GroupID: "cloud-watch-worker",
	})

	defer reader.Close()

	log.Println("Worker started...")

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		var event TelemetryEvent

		err = json.Unmarshal(message.Value, &event)
		if err != nil {
			log.Println("Invalid event:", err)
			continue
		}

		fmt.Printf(
			"Received event: server=%s cpu=%.2f memory=%.2f latency=%.2f status=%s\n",
			event.ServerID,
			event.CPU,
			event.Memory,
			event.Latency,
			event.Status,
		)
	}
}