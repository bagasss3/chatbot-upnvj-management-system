package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type fallbackChatLogService struct {
	fallbackChatLogRepository model.FallbackChatLogRepository
}

func NewFallbackChatLogService(fallbackChatLogRepository model.FallbackChatLogRepository) model.FallbackChatLogService {
	return &fallbackChatLogService{
		fallbackChatLogRepository: fallbackChatLogRepository,
	}
}

func (fcl *fallbackChatLogService) CreateFallbackChatLog(ctx context.Context, req model.CreateFallbackChatLogRequest) (*model.FallbackChatLog, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	newFallbackChatLogIntent := &model.FallbackChatLog{
		Id:   helper.GenerateID(),
		Chat: req.Chat,
	}

	err := fcl.fallbackChatLogRepository.Create(ctx, newFallbackChatLogIntent)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return newFallbackChatLogIntent, nil
}

func (fcl *fallbackChatLogService) FindAllFallbackChatLog(ctx context.Context, page string) ([]*model.FallbackChatLog, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	if page != string(model.FallbackPageDashboard) && page != string(model.FallbackPageLog) {
		log.Error(constant.ErrInvalidArgument)
		return nil, constant.ErrInvalidArgument
	}

	fallbacks, err := fcl.fallbackChatLogRepository.FindAll(ctx, page)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return fallbacks, nil
}
