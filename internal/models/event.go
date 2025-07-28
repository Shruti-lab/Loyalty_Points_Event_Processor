package models

import "time"

type LoyaltyEvent struct {
	UserID    string    `json:"user_id"`
	EventType string    `json:"event_type"` // purchase, reward, redeem
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}
