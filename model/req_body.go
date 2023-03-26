package model

import (
	"time"
)

type ReqBody struct {
	Id           int64     `json:"id"`
	ActionHttpId int64     `json:"action_http_id"`
	ReqName      string    `json:"req_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
