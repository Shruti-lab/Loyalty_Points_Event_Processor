# Loyalty Points Event Processor

Event-driven platform simulating loyalty reward events using Go + Kafka + Postgres.

## Tech Stack
- Golang (producer/consumer)
- Apache Kafka (event bus)
- Postgres (state store)
- Docker (infra)

## Quick Summary of Components

| Component            | Description                                                          |
| -------------------- | -------------------------------------------------------------------- |
| `producer/`          | Sends simulated `purchase` / `redeem` events to Kafka.               |
| `consumer/`          | Reads from Kafka, applies rules, updates Postgres points table.      |
| `api/`               | REST service with `GET /points/{user_id}` and Prometheus `/metrics`. |
| `prometheus.yml`     | Scrapes `/metrics` endpoint to collect API stats.                    |
| `docker-compose.yml` | Starts Kafka, Postgres, Prometheus, Grafana.                         |
| `.github/workflows/` | GitHub Actions: lint, build, test on push/PR.                        |
| `config/rules.yaml`  | (Optional) External rule definitions — upcoming.                     |


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

# Terminal 3: Start API (includes metrics on :8080/metrics)
go run cmd/api/main.go

# Test API:
curl http://localhost:8080/points/user-123

# Open Prometheus UI
http://localhost:9090

# Open Grafana (login: admin / admin)
http://localhost:3000


```
## Project Structure
```
loyalty-points-system/
│
├── cmd/                            # Main entrypoints for each service
│   ├── producer/                   # Kafka event simulator (writes to Kafka)
│   │   └── main.go
│   ├── consumer/                   # Kafka consumer + points processor
│   │   └── main.go
│   └── api/                        # REST API for querying user points
│       └── main.go
│
├── internal/
│   ├── db/                         # Postgres interaction layer
│   │   └── postgres.go
│   └── kafka/                      # Kafka setup
│       ├── writer.go
│
├── models/                         # Event & DB struct definitions
│   └── event.go
|
├── monitoring/                    # Prometheus config
│   └── prometheus.yml
│
├── .github/
│   └── workflows/
│       └── go-ci.yml              # GitHub Actions CI/CD pipeline
│
├── Dockerfile                     # Build image for API or all-in-one
├── docker-compose.yml             # Run Kafka, Postgres, Prometheus, Grafana
├── go.mod / go.sum                # Dependencies
└── README.md                      # Project overview

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

    
    
