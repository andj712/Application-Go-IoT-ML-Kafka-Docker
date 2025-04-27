package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type DeviceData struct {
	DeviceID    string  `json:"device_id"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func main() {
	// Kafka reader konfiguracija
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "iot-topic",
		GroupID: "iot-backend-group",
	})
	defer reader.Close()

	fmt.Println("Backend Kafka reader started, listening on topic iot-topic...")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("ReadMessage error:", err)
			continue
		}

		var data DeviceData
		if err := json.Unmarshal(m.Value, &data); err != nil {
			log.Println("Unmarshal error:", err)
			continue
		}

		fmt.Printf("Received from Kafka: %+v\n", data)
	}
}
