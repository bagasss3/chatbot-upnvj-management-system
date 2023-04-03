package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type intentService struct {
	intentRepository model.IntentRepository
}

func NewIntentService(intentRepository model.IntentRepository) model.IntentService {
	return &intentService{
		intentRepository: intentRepository,
	}
}

func (i *intentService) CreateIntent(ctx context.Context, req model.CreateUpdateIntentRequest) (*model.Intent, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	intent := &model.Intent{
		Id:   helper.GenerateID(),
		Name: req.Name,
	}

	err := i.intentRepository.Create(ctx, intent)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return intent, nil
}

func (i *intentService) FindAllIntent(ctx context.Context) ([]*model.Intent, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	intents, err := i.intentRepository.FindAll(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return intents, nil
}

func (i *intentService) FindIntentByID(ctx context.Context, id int64) (*model.Intent, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	if id <= 0 {
		return nil, constant.ErrInvalidArgument
	}

	intent, err := i.intentRepository.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if intent == nil {
		return nil, constant.ErrNotFound
	}

	return intent, nil
}

func (i *intentService) UpdateIntent(ctx context.Context, id int64, req model.CreateUpdateIntentRequest) (*model.Intent, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
		"req": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	intent, err := i.FindIntentByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	intent.Name = req.Name
	err = i.intentRepository.Update(ctx, intent.Id, intent)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return intent, nil
}

func (i *intentService) DeleteIntent(ctx context.Context, id int64) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	_, err := i.FindIntentByID(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	err = i.intentRepository.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}
