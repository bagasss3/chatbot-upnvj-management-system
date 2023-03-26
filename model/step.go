package model

import (
	"time"
)

type StepType string

const (
	StepIntent    StepType = "INTENT"
	StepUtterance StepType = "UTTERANCE"
	StepAction    StepType = "ACTION"
)

type Step struct {
	Id        int64     `json:"id"`
	StoryId   int64     `json:"story_id"`
	Type      StepType  `json:"type"`
	DataId    int64     `json:"data_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
