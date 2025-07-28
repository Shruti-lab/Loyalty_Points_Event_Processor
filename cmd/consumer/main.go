package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/yourname/loyalty-platform/internal/db"
	"github.com/yourname/loyalty-platform/internal/models"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "loyalty-consumers",
		Topic:   "loyalty-events",
	})
	defer reader.Close()

	conn := db.InitDB()
	defer conn.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		var e models.LoyaltyEvent
		if err := json.Unmarshal(m.Value, &e); err != nil {
			log.Println("Unmarshal error:", err)
			continue
		}
		fmt.Println("Consumed:", e)
			err = db.ProcessEvent(conn, e)
				if err != nil {
					log.Println("Process error:", err)
				}

	}
}
