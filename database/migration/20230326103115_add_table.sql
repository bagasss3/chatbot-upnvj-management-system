-- +goose Up

CREATE TABLE IF NOT EXISTS users (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(100),
  email VARCHAR(100),
  password TEXT,
  type ENUM('ADMIN', 'SUPER_ADMIN'),
  major_id VARCHAR(100),
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

CREATE TABLE IF NOT EXISTS sessions (
  id VARCHAR(100) PRIMARY KEY,
  access_token TEXT NOT NULL,
  refresh_token TEXT NOT NULL,
  access_token_expired_at timestamp NOT NULL DEFAULT NOW(),
  refresh_token_expired_at timestamp NOT NULL DEFAULT NOW(),
  user_id VARCHAR(100) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

ALTER TABLE sessions ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS intents (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

CREATE TABLE IF NOT EXISTS utterances (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  response TEXT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

CREATE TABLE IF NOT EXISTS stories (
  id VARCHAR(100) PRIMARY KEY,
  story_title TEXT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

CREATE TABLE IF NOT EXISTS entities (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  intent_id VARCHAR(100) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

ALTER TABLE entities ADD FOREIGN KEY (intent_id) REFERENCES intents(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS examples (
  id VARCHAR(100) PRIMARY KEY,
  example TEXT NOT NULL,
  intent_id VARCHAR(100) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

ALTER TABLE examples ADD FOREIGN KEY (intent_id) REFERENCES intents(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS rules (
  id VARCHAR(100) PRIMARY KEY,
  rule_title TEXT NOT NULL,
  intent_id VARCHAR(100) NOT NULL,
  response_id VARCHAR(100) NOT NULL,
  type ENUM('UTTERANCE', 'ACTION_HTTP'),
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

ALTER TABLE rules ADD FOREIGN KEY (intent_id) REFERENCES intents(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS action_https (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
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

CREATE TABLE IF NOT EXISTS krs_actions (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  get_http_req TEXT NOT NULL,
  api_key TEXT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

CREATE TABLE IF NOT EXISTS req_bodies (
  id VARCHAR(100) PRIMARY KEY,
  action_http_id VARCHAR(100) NOT NULL,
  req_name VARCHAR(150) NOT NULL,
  data_type ENUM('STRING', 'INT','FLOAT','DATE'),
  method ENUM('POST','PUT'),
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

ALTER TABLE req_bodies ADD FOREIGN KEY (action_http_id) REFERENCES action_https(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS steps (
  id VARCHAR(100) PRIMARY KEY,
  story_id VARCHAR(100) NOT NULL,
  type ENUM('INTENT', 'UTTERANCE','ACTION_HTTP'),
  response_id VARCHAR(100) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

ALTER TABLE steps ADD FOREIGN KEY (story_id) REFERENCES stories(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS log_intents (
  id VARCHAR(100) PRIMARY KEY,
  intent_id VARCHAR(100) NOT NULL,
  mention INT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

ALTER TABLE log_intents ADD FOREIGN KEY (intent_id) REFERENCES intents(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS configurations (
  id VARCHAR(100) PRIMARY KEY,
  diet_classifier_epoch INT NOT NULL,
  fallback_classifier_treshold FLOAT NOT NULL,
  response_selector_epoch INT NOT NULL,
  ted_policy_epoch INT NOT NULL,
  fallback_utterance_id VARCHAR(100) NOT NULL,
  fallback_treshold FLOAT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

ALTER TABLE configurations ADD FOREIGN KEY (fallback_utterance_id) REFERENCES utterances(id) ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS training_histories (
  id VARCHAR(100) PRIMARY KEY,
  user_id VARCHAR(100) NOT NULL,
  total_time INT NOT NULL,
  status ENUM('DONE', 'FAILED'),
  created_at timestamp NOT NULL DEFAULT NOW()
);

ALTER TABLE training_histories ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE;

-- +goose Down

DROP TABLE IF EXISTS training_histories;
DROP TABLE IF EXISTS configurations;
DROP TABLE IF EXISTS log_intents;
DROP TABLE IF EXISTS steps;
DROP TABLE IF EXISTS req_bodies;
DROP TABLE IF EXISTS krs_actions;
DROP TABLE IF EXISTS action_https;
DROP TABLE IF EXISTS rules;
DROP TABLE IF EXISTS examples;
DROP TABLE IF EXISTS entities;
DROP TABLE IF EXISTS stories;
DROP TABLE IF EXISTS utterances;
DROP TABLE IF EXISTS intents;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS users;
