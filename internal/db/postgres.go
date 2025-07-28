package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/Shruti-lab/Loyalty_Points_Event_Processor/internal/models"
)

func InitDB() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=user password=pass dbname=loyalty sslmode=disable")
	if err != nil {
		log.Fatal("DB connect error:", err)
	}
	return db
}

func ProcessEvent(db *sql.DB, e models.LoyaltyEvent) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// insert event log
	_, err = tx.Exec(`INSERT INTO loyalty_events (user_id, event_type, amount, timestamp) VALUES ($1, $2, $3, $4)`,
		e.UserID, e.EventType, e.Amount, e.Timestamp)
	if err != nil {
		return err
	}

	// point calculation logic
	var pointsDelta int
	switch e.EventType {
	case "purchase":
		pointsPerDollar := 1
		if e.Timestamp.Weekday() == time.Saturday || e.Timestamp.Weekday() == time.Sunday {
			pointsPerDollar = 2
		}
		pointsDelta = int(e.Amount) * pointsPerDollar
	case "redeem":
		pointsDelta = -int(e.Amount) // amount == points to deduct
	default:
		pointsDelta = 0
	}

	// upsert user points
	_, err = tx.Exec(`
		INSERT INTO user_points (user_id, points)
		VALUES ($1, $2)
		ON CONFLICT (user_id)
		DO UPDATE SET points = user_points.points + $2`, e.UserID, pointsDelta)
	if err != nil {
		return err
	}

	return tx.Commit()
}

