-- +goose Up
ALTER TABLE fallback_chat_logs
ADD COLUMN cluster INT;

-- +goose Down
ALTER TABLE fallback_chat_logs
DROP COLUMN IF EXISTS cluster;
