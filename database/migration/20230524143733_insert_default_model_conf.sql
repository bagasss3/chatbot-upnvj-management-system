-- +goose Up
INSERT INTO configurations (id,diet_classifier_epoch,fallback_classifier_treshold,response_selector_epoch,ted_policy_epoch,fallback_utterance_id,fallback_treshold,unexpected_intent_policy_epoch) VALUES ("1",100,0.3,100,100,"1",0.5,0);


-- +goose Down
DELETE FROM configurations WHERE id IN ("1");

