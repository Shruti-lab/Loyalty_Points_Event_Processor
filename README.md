# Loyalty Points Event Processor

Event-driven platform simulating loyalty reward events using Go + Kafka + Postgres.

## Tech Stack
- Golang (producer/consumer)
- Apache Kafka (event bus)
- Postgres (state store)
- Docker (infra)

## Features
- Kafka Producer simulates `purchase` events
- Kafka Consumer logs and persists events to DB
- Easily extendable to support reward/redemption logic

## Run It
```bash
# Start containers
docker-compose up -d

# Terminal 1: Produce events
go run cmd/producer/main.go

# Terminal 2: Consume + store events + apply logic
go run cmd/consumer/main.go

# Terminal 3: Start API
go run cmd/api/main.go

# Test API:
curl http://localhost:8080/points/user-123

```
## Project Structure

loyalty-platform/
├── docker-compose.yml
├── go.mod
├── go.sum
├── README.md
├── cmd/
│   ├── producer/           # Kafka producer: simulates events
│   │   └── main.go
│   └── consumer/           # Kafka consumer: stores in DB
│       └── main.go
├── internal/
│   ├── kafka/
│   │   └── writer.go       # Kafka writer setup
│   ├── models/
│   │   └── event.go        # Event schema
│   └── db/
│       └── postgres.go     # DB insert logic
└── sql/
    └── init.sql            # Postgres schema
