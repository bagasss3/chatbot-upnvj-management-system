-- +goose Up
INSERT INTO utterances (id, name, response) VALUES ("1", 'utter_fallback', 'Saya tidak mengerti maksud perkataan anda. Coba parafrase kalimat anda');


-- +goose Down
DELETE FROM utterances WHERE id IN ("1");
