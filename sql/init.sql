CREATE TABLE IF NOT EXISTS loyalty_events (
  id SERIAL PRIMARY KEY,
  user_id TEXT NOT NULL,
  event_type TEXT NOT NULL,
  amount NUMERIC,
  timestamp TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_points (
  user_id TEXT PRIMARY KEY,
  points INTEGER DEFAULT 0
);
