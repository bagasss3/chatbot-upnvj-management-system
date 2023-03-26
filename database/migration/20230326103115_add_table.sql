-- +goose Up

BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS user (
  id BIGINT PRIMARY KEY,
  name VARCHAR(100),
  email VARCHAR(100),
  password VARCHAR(50),
  type ENUM('ADMIN', 'SUPER_ADMIN'),
  major_id BIGINT,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

CREATE TABLE IF NOT EXISTS session (
  id BIGINT PRIMARY KEY,
  access_token TEXT NOT NULL,
  refresh_token TEXT NOT NULL,
  refresh_token_expired_at timestamp NOT NULL,
  user_id BIGINT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

ALTER TABLE session ADD FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS intent (
  id BIGINT PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

CREATE TABLE IF NOT EXISTS utterance (
  id BIGINT PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  response TEXT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

CREATE TABLE IF NOT EXISTS action (
  id BIGINT PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

CREATE TABLE IF NOT EXISTS story (
  id BIGINT PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

CREATE TABLE IF NOT EXISTS entity (
  id BIGINT PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  intent_id BIGINT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

ALTER TABLE entity ADD FOREIGN KEY (intent_id) REFERENCES intent(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS example (
  id BIGINT PRIMARY KEY,
  example TEXT NOT NULL,
  intent_id BIGINT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

ALTER TABLE example ADD FOREIGN KEY (intent_id) REFERENCES intent(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS rule (
  id BIGINT PRIMARY KEY,
  example TEXT NOT NULL,
  intent_id BIGINT NOT NULL,
  response_id BIGINT NOT NULL,
  type ENUM('UTTERANCE', 'ACTION'),
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

ALTER TABLE rule ADD FOREIGN KEY (intent_id) REFERENCES intent(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS action_http (
  id BIGINT PRIMARY KEY,
  action_id BIGINT NOT NULL,
  get_http_req TEXT NOT NULL,
  post_http_req TEXT NOT NULL,
  put_http_req TEXT NOT NULL,
  del_http_req TEXT NOT NULL,
  api_key TEXT NOT NULL,
  text_response TEXT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

ALTER TABLE action_http ADD FOREIGN KEY (action_id) REFERENCES action(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS krs_action (
  id BIGINT PRIMARY KEY,
  action_id BIGINT NOT NULL,
  get_http_req TEXT NOT NULL,
  api_key TEXT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

ALTER TABLE krs_action ADD FOREIGN KEY (action_id) REFERENCES action(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS req_body (
  id BIGINT PRIMARY KEY,
  action_http_id BIGINT NOT NULL,
  req_name VARCHAR(150) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

ALTER TABLE req_body ADD FOREIGN KEY (action_http_id) REFERENCES action_http(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS step (
  id BIGINT PRIMARY KEY,
  story_id BIGINT NOT NULL,
  type ENUM('INTENT', 'UTTERANCE','ACTION'),
  response_id BIGINT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

ALTER TABLE step ADD FOREIGN KEY (story_id) REFERENCES story(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS log_intent (
  id BIGINT PRIMARY KEY,
  intent_id BIGINT NOT NULL,
  mention INT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

ALTER TABLE log_intent ADD FOREIGN KEY (intent_id) REFERENCES intent(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS configuration (
  id BIGINT PRIMARY KEY,
  dietclassifier_epoch INT NOT NULL,
  fallbackclassifier_treshold INT NOT NULL,
  responseselector_epoch INT NOT NULL,
  tedpolicy_epoch INT NOT NULL,
  fallback_utterance_id BIGINT NOT NULL,
  fallback_treshold INT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

ALTER TABLE configuration ADD FOREIGN KEY (fallback_utterance_id) REFERENCES utterance(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS training_history (
  id BIGINT PRIMARY KEY,
  user_id BIGINT NOT NULL,
  total_time INT NOT NULL,
  status ENUM('DONE', 'FAILED'),
  created_at timestamp NOT NULL DEFAULT NOW()
);

ALTER TABLE training_history ADD FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE RESTRICT ON UPDATE CASCADE;

COMMIT;

-- +goose Down

BEGIN TRANSACTION;

DROP TABLE IF EXISTS training_history;
DROP TABLE IF EXISTS configuration;
DROP TABLE IF EXISTS log_intent;
DROP TABLE IF EXISTS step;
DROP TABLE IF EXISTS req_body;
DROP TABLE IF EXISTS krs_action;
DROP TABLE IF EXISTS action_http;
DROP TABLE IF EXISTS rule;
DROP TABLE IF EXISTS example;
DROP TABLE IF EXISTS entity;
DROP TABLE IF EXISTS story;
DROP TABLE IF EXISTS action;
DROP TABLE IF EXISTS utterance;
DROP TABLE IF EXISTS intent;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS users;

COMMIT;