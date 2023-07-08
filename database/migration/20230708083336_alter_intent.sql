-- +goose Up
ALTER TABLE intents
ADD COLUMN is_information_academic BOOLEAN DEFAULT false;

-- +goose Down
ALTER TABLE intents
DROP COLUMN IF EXISTS is_information_academic;
