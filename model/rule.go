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
	IntentId   string   `json:"intent_id" validate:"required"`
	ResponseId string   `json:"response_id" validate:"required"`
	RuleTitle  string   `json:"rule_title" validate:"required,min=3,max=120"`
	Type       RuleType `json:"type" validate:"required"`
}

func (c *CreateUpdateRuleRequest) Validate() error {
	return validate.Struct(c)
}

type Rule struct {
	Id         string         `json:"id"`
	IntentId   string         `json:"intent_id"`
	ResponseId string         `json:"response_id"`
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
	FindAllRule(ctx context.Context, name string) ([]*Rule, error)
	FindRuleByID(ctx context.Context, id string) (*Rule, error)
	UpdateRule(ctx context.Context, id string, req CreateUpdateRuleRequest) (*Rule, error)
	DeleteRule(ctx context.Context, id string) (bool, error)
}

type RuleRepository interface {
	Create(ctx context.Context, rule *Rule) error
	FindAll(ctx context.Context, name string) ([]*Rule, error)
	FindByID(ctx context.Context, id string) (*Rule, error)
	Update(ctx context.Context, id string, rule *Rule) error
	Delete(ctx context.Context, id string) error
	CountAll(ctx context.Context) (int64, error)
}
