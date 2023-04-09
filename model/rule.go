package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RuleType string

const (
	RuleUtterance RuleType = "UTTERANCE"
	RuleAction    RuleType = "ACTION_HTTP"
)

type CreateUpdateRuleRequest struct {
	IntentId   int64    `json:"intent_id" validate:"required"`
	ResponseId int64    `json:"response_id" validate:"required"`
	RuleTitle  string   `json:"rule_title" validate:"required,min=3,max=60"`
	Type       RuleType `json:"type" validate:"required"`
}

func (c *CreateUpdateRuleRequest) Validate() error {
	return validate.Struct(c)
}

type Rule struct {
	Id         int64          `json:"id"`
	IntentId   int64          `json:"intent_id"`
	ResponseId int64          `json:"response_id"`
	RuleTitle  string         `json:"rule_title"`
	Type       RuleType       `json:"type"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type RuleController interface {
	HandleCreateRule() echo.HandlerFunc
	HandleFindAllRule() echo.HandlerFunc
	HandleFindRuleByID() echo.HandlerFunc
	HandleUpdateRule() echo.HandlerFunc
	HandleDeleteRule() echo.HandlerFunc
}

type RuleService interface {
	CreateRule(ctx context.Context, req CreateUpdateRuleRequest) (*Rule, error)
	FindAllRule(ctx context.Context) ([]*Rule, error)
	FindRuleByID(ctx context.Context, id int64) (*Rule, error)
	UpdateRule(ctx context.Context, id int64, req CreateUpdateRuleRequest) (*Rule, error)
	DeleteRule(ctx context.Context, id int64) (bool, error)
}

type RuleRepository interface {
	Create(ctx context.Context, rule *Rule) error
	FindAll(ctx context.Context) ([]*Rule, error)
	FindByID(ctx context.Context, id int64) (*Rule, error)
	Update(ctx context.Context, id int64, rule *Rule) error
	Delete(ctx context.Context, id int64) error
}
