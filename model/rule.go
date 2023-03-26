package model

import (
	"time"

	"gorm.io/gorm"
)

type RuleType string

const (
	RuleUtterance RuleType = "UTTERANCE"
	RuleAction    RuleType = "ACTION"
)

type Rule struct {
	Id        int64          `json:"id"`
	IntentId  int64          `json:"intent_id"`
	DataId    int64          `json:"data_id"`
	RuleTitle string         `json:"rule_title"`
	Type      RuleType       `json:"type"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
