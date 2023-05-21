package service

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type conversationService struct {
	ruleRepository  model.RuleRepository
	storyRepository model.StoryRepository
}

func NewConversationService(ruleRepository model.RuleRepository, storyRepository model.StoryRepository) model.ConversationService {
	return &conversationService{
		ruleRepository:  ruleRepository,
		storyRepository: storyRepository,
	}
}

func (c *conversationService) CountAllConversation(ctx context.Context) (int64, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	countStory, err := c.storyRepository.CountAll(ctx)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	countRule, err := c.ruleRepository.CountAll(ctx)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return countRule + countStory, nil
}
