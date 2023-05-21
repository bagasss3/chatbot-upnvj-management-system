package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type logIntentService struct {
	logIntentRepository model.LogIntentRepository
	intentRepository    model.IntentRepository
}

func NewLogIntentService(logIntentRepository model.LogIntentRepository, intentRepository model.IntentRepository) model.LogIntentService {
	return &logIntentService{
		logIntentRepository: logIntentRepository,
		intentRepository:    intentRepository,
	}
}

func (li *logIntentService) CreateOrUpdateLogIntent(ctx context.Context, req model.CreateUpdateLogIntentRequest) (*model.LogIntent, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	intent, err := li.intentRepository.FindByID(ctx, req.IntentId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if intent == nil {
		log.Error("intent not found")
		return nil, constant.ErrNotFound
	}

	logIntent, err := li.FindLogIntentByIntentID(ctx, intent.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if logIntent != nil {
		logIntent.Mention++
		err := li.logIntentRepository.Update(ctx, logIntent.IntentId, logIntent)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		logIntent.Intent = *intent
		return logIntent, nil
	}

	newLogIntent := &model.LogIntent{
		Id:       helper.GenerateID(),
		IntentId: req.IntentId,
		Mention:  1,
		Intent:   *intent,
	}

	err = li.logIntentRepository.Create(ctx, newLogIntent)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return newLogIntent, nil
}

func (li *logIntentService) FindAllLogIntent(ctx context.Context) ([]*model.LogIntent, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	logIntents, err := li.logIntentRepository.FindAll(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return logIntents, nil
}

func (li *logIntentService) FindLogIntentByIntentID(ctx context.Context, intentId string) (*model.LogIntent, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	if intentId == "" {
		return nil, constant.ErrInvalidArgument
	}

	logIntent, err := li.logIntentRepository.FindByIntentID(ctx, intentId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return logIntent, nil
}
