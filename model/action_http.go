package model

import (
	"time"

	"gorm.io/gorm"
)

type ActionHttp struct {
	Id           int64          `json:"id"`
	ActionId     int64          `json:"action_id"`
	GetHttpReq   string         `json:"get_http_req"`
	PostHttpReq  string         `json:"post_http_req"`
	PutHttpReq   string         `json:"put_http_req"`
	DelHttpReq   string         `json:"del_http_req"`
	ApiKey       string         `json:"api_key"`
	TextResponse string         `json:"text_response"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
