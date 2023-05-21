-- +goose Up
CREATE TABLE IF NOT EXISTS fallback_chat_logs (
  id VARCHAR(100) PRIMARY KEY,
  chat TEXT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

-- +goose Down
DROP TABLE IF EXISTS fallback_chat_logs;

