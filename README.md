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
```
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
```

**Description:**
- **cmd/producer/**: Contains the main entry point for the Kafka event producer.
- **cmd/consumer/**: Contains the main entry point for the Kafka event consumer, which processes and stores events in the database.
- **internal/kafka/**: Kafka writer setup and configuration.
- **internal/models/**: Go structures defining the event schema.
- **internal/db/**: Database interaction logic, specifically for PostgreSQL.
- **sql/**: SQL scripts for initializing the database schema.
- **docker-compose.yml**: Orchestrates services for local development.
- **go.mod / go.sum**: Manage Go modules and dependencies.

    
    
