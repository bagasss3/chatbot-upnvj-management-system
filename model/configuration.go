package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

type CreateUpdateConfigurationRequest struct {
	DietClassifierEpoch        int32   `json:"diet_classifier_epoch" validate:"required"`
	FallbackClassifierTreshold float32 `json:"fallback_classifier_treshold" validate:"required"`
	ResponseSelectorEpoch      int32   `json:"response_selector_epoch" validate:"required"`
	TedPolicyEpoch             int32   `json:"ted_policy_epoch" validate:"required"`
	FallbackUtteranceId        string  `json:"fallback_utterance_id" validate:"required"`
	FallbackTreshold           float32 `json:"fallback_treshold" validate:"required"`
}

func (c *CreateUpdateConfigurationRequest) Validate() error {
	return validate.Struct(c)
}

type Configuration struct {
	Id                         string    `json:"id"`
	DietClassifierEpoch        int32     `json:"diet_classifier_epoch"`
	FallbackClassifierTreshold float32   `json:"fallback_classifier_treshold"`
	ResponseSelectorEpoch      int32     `json:"response_selector_epoch"`
	TedPolicyEpoch             int32     `json:"ted_policy_epoch"`
	FallbackUtteranceId        string    `json:"fallback_utterance_id"`
	FallbackTreshold           float32   `json:"fallback_treshold"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}

type ConfigurationController interface {
	HandleCreateConfiguration() echo.HandlerFunc
	HandleFindAllConfiguration() echo.HandlerFunc
	HandleFindConfiguration() echo.HandlerFunc
	HandleUpdateConfiguration() echo.HandlerFunc
}

type ConfigurationService interface {
	CreateConfiguration(ctx context.Context, req CreateUpdateConfigurationRequest) (*Configuration, error)
	FindAllConfiguration(ctx context.Context) ([]*Configuration, error)
	FindConfiguration(ctx context.Context, id string) (*Configuration, error)
	UpdateConfiguration(ctx context.Context, id string, req CreateUpdateConfigurationRequest) (*Configuration, error)
}

type ConfigurationRepository interface {
	Create(ctx context.Context, configuration *Configuration) error
	FindAll(ctx context.Context) ([]*Configuration, error)
	FindByID(ctx context.Context, id string) (*Configuration, error)
	Update(ctx context.Context, id string, configuration *Configuration) error
}
