package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

type DeviceData struct {
	DeviceID    string  `json:"device_id"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func generateRandomData(deviceID string) DeviceData {
	return DeviceData{
		DeviceID:    deviceID,
		Temperature: 20.0 + rand.Float64()*15.0,
		Humidity:    30.0 + rand.Float64()*40.0,
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Kafka writer konfiguracija
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "iot-topic",
	})
	defer writer.Close()

	deviceIDs := []string{"sensor-1", "sensor-2", "sensor-3"}

	for {
		for _, id := range deviceIDs {
			data := generateRandomData(id)
			// serijalizacija u JSON
			value, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Marshal error:", err)
				continue
			}

			msg := kafka.Message{
				Key:   []byte(data.DeviceID),
				Value: value,
			}

			if err := writer.WriteMessages(context.Background(), msg); err != nil {
				fmt.Println("WriteMessages error:", err)
			} else {
				fmt.Printf("Sent to Kafka: %+v\n", data)
			}
		}
		time.Sleep(5 * time.Second)
	}
}
