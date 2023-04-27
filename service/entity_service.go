package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type entityService struct {
	entityRepository model.EntityRepository
	intentRepository model.IntentRepository
}

func NewEntityService(entityRepository model.EntityRepository, intentRepository model.IntentRepository) model.EntityService {
	return &entityService{
		entityRepository: entityRepository,
		intentRepository: intentRepository,
	}
}

func (e *entityService) CreateEntity(ctx context.Context, req model.CreateEntityRequest) (*model.Entity, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	intent, err := e.intentRepository.FindByID(ctx, req.IntentId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if intent == nil {
		log.Error(constant.ErrNotFound)
		return nil, constant.ErrNotFound
	}

	entity := &model.Entity{
		Id:       helper.GenerateID(),
		Name:     req.Name,
		IntentId: req.IntentId,
	}

	err = e.entityRepository.Create(ctx, entity)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return entity, nil
}

func (e *entityService) FindAllEntity(ctx context.Context, intentId string) ([]*model.Entity, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	entities, err := e.entityRepository.FindAll(ctx, intentId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return entities, nil
}

func (e *entityService) FindEntityByID(ctx context.Context, id, intentId string) (*model.Entity, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	if id == "" || intentId == "" {
		return nil, constant.ErrInvalidArgument
	}

	intent, err := e.intentRepository.FindByID(ctx, intentId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if intent == nil {
		return nil, constant.ErrNotFound
	}

	entity, err := e.entityRepository.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if entity == nil {
		return nil, constant.ErrNotFound
	}

	return entity, nil
}

func (e *entityService) DeleteEntity(ctx context.Context, id, intentId string) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	_, err := e.FindEntityByID(ctx, id, intentId)
	if err != nil {
		log.Error(err)
		return false, err
	}

	err = e.entityRepository.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}
