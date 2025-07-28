package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/Shruti-lab/Loyalty_Points_Event_Processor/internal/models"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "loyalty-events",
	})
	defer writer.Close()

	users := []string{"user-123", "user-456", "user-789"}

	for {
		e := models.LoyaltyEvent{
			UserID:    users[rand.Intn(len(users))],
			EventType: "purchase",
			Amount:    float64(rand.Intn(500)),
			Timestamp: time.Now(),
		}
		payload, _ := json.Marshal(e)
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(e.UserID),
				Value: payload,
			})
		if err != nil {
			fmt.Println("Failed to write message:", err)
		} else {
			fmt.Println("Produced:", string(payload))
		}
		time.Sleep(2 * time.Second)
	}
}
