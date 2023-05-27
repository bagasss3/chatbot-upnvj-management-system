-- +goose Up
INSERT INTO intents (id,name) VALUES ("1","nlu_fallback");
INSERT INTO rules (id,rule_title,intent_id,response_id,type) VALUES ("1","Ask the user to rephrase whenever they send a message with low NLU confidence","1","1","UTTERANCE");

-- +goose Down
DELETE FROM rules WHERE id IN ("1");
DELETE FROM intents WHERE id IN ("1");

