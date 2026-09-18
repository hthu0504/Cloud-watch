package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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
	// -----------------------------
	// AWS / DynamoDB setup
	// -----------------------------

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("us-east-2"),
	)
	if err != nil {
		log.Fatal("Failed to load AWS config:", err)
	}

	dynamoClient := dynamodb.NewFromConfig(cfg)

	tableName := "CloudwatchServer"

	// -----------------------------
	// Kafka setup
	// -----------------------------

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "telemetry",
		GroupID: "cloud-watch-worker",
	})

	defer reader.Close()

	log.Println("Worker started...")

	// -----------------------------
	// Main worker loop
	// -----------------------------

	for {
		message, err := reader.ReadMessage(ctx)
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

		// -----------------------------
		// Save server state to DynamoDB
		// -----------------------------

		_, err = dynamoClient.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: &tableName,
			Item: map[string]types.AttributeValue{
				"server_id": &types.AttributeValueMemberS{
					Value: event.ServerID,
				},
				"cpu": &types.AttributeValueMemberN{
					Value: fmt.Sprintf("%.2f", event.CPU),
				},
				"memory": &types.AttributeValueMemberN{
					Value: fmt.Sprintf("%.2f", event.Memory),
				},
				"latency": &types.AttributeValueMemberN{
					Value: fmt.Sprintf("%.2f", event.Latency),
				},
				"status": &types.AttributeValueMemberS{
					Value: event.Status,
				},
			},
		})

		if err != nil {
			log.Println("Failed to update DynamoDB:", err)
			continue
		}

		log.Printf("DynamoDB updated: server=%s\n", event.ServerID)
	}
}
