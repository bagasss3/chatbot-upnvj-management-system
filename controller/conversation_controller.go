package controller

import (
	"cbupnvj/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type conversationController struct {
	conversationService model.ConversationService
}

func NewConversationController(conversationService model.ConversationService) model.ConversationController {
	return &conversationController{
		conversationService: conversationService,
	}
}

func (co *conversationController) HandleCountAllConversation() echo.HandlerFunc {
	return func(c echo.Context) error {
		count, err := co.conversationService.CountAllConversation(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, count)
	}
}
