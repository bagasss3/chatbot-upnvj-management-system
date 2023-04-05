package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type actionService struct {
	actionRepository model.ActionRepository
}

func NewActionService(actionRepository model.ActionRepository) model.ActionService {
	return &actionService{
		actionRepository: actionRepository,
	}
}

func (a *actionService) CreateAction(ctx context.Context, req model.CreateUpdateActionRequest) (*model.Action, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	action := &model.Action{
		Id:   helper.GenerateID(),
		Name: req.Name,
	}

	err := a.actionRepository.Create(ctx, action)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return action, nil
}

func (a *actionService) FindAllAction(ctx context.Context) ([]*model.Action, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	actions, err := a.actionRepository.FindAll(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return actions, nil
}

func (a *actionService) FindActionByID(ctx context.Context, id int64) (*model.Action, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	if id <= 0 {
		return nil, constant.ErrInvalidArgument
	}

	action, err := a.actionRepository.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if action == nil {
		return nil, constant.ErrNotFound
	}

	return action, nil
}

func (a *actionService) UpdateAction(ctx context.Context, id int64, req model.CreateUpdateActionRequest) (*model.Action, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
		"req": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	action, err := a.FindActionByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	action.Name = req.Name
	err = a.actionRepository.Update(ctx, action.Id, action)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return action, nil
}

func (a *actionService) DeleteAction(ctx context.Context, id int64) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	_, err := a.FindActionByID(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	err = a.actionRepository.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}
