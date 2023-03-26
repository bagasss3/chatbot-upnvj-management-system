package model

import "time"

type Example struct {
	Id        int64     `json:"id"`
	IntentId  int64     `json:"intent_id"`
	Example   string    `json:"example"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
