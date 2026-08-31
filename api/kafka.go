package main

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

var kafkaWriter = &kafka.Writer{
	Addr:     kafka.TCP("localhost:9092"),
	Topic:    "telemetry",
	Balancer: &kafka.LeastBytes{},
}

func publishEvent(event interface{}) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return kafkaWriter.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: data,
		},
	)
}