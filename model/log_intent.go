package model

import "time"

type LogIntent struct {
	Id        int64     `json:"id"`
	IntentId  int64     `json:"intent_id"`
	Mention   int       `json:"mention"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
