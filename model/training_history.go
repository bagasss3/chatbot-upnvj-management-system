package model

import "time"

type Status string

const (
	StatusDone   Status = "DONE"
	StatusFailed Status = "FAILED"
)

type TrainingHistory struct {
	Id        int64     `json:"int64"`
	UserId    int64     `json:"user_id"`
	TotalTime int       `json:"total_time"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
