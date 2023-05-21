package model

import (
	"context"

	"github.com/labstack/echo/v4"
)

type ConversationController interface {
	HandleCountAllConversation() echo.HandlerFunc
}

type ConversationService interface {
	CountAllConversation(ctx context.Context) (int64, error)
}
