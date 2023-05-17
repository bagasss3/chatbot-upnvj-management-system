-- +goose Up
ALTER TABLE configurations
ADD COLUMN unexpected_intent_policy_epoch INT NOT NULL AFTER ted_policy_epoch;

-- +goose Down
ALTER TABLE configurations
DROP COLUMN unexpected_intent_policy_epoch;
