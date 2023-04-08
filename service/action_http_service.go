package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type actionHttpService struct {
	actionHttpRepository model.ActionHttpRepository
	actionRepository     model.ActionRepository
}

func NewActionHttpService(actionHttpRepository model.ActionHttpRepository, actionRepository model.ActionRepository) model.ActionHttpService {
	return &actionHttpService{
		actionHttpRepository: actionHttpRepository,
		actionRepository:     actionRepository,
	}
}

func (a *actionHttpService) CreateActionHttp(ctx context.Context, req model.CreateActionHttpRequest) (*model.ActionHttp, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	action, err := a.actionRepository.FindByID(ctx, req.ActionId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if action == nil {
		log.Error(constant.ErrNotFound)
		return nil, constant.ErrNotFound
	}

	isActionUsed, err := a.actionHttpRepository.FindByActionID(ctx, req.ActionId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if isActionUsed != nil {
		log.Error(constant.ErrAlreadyExists)
		return nil, constant.ErrAlreadyExists
	}

	actionHttp := &model.ActionHttp{
		Id:           helper.GenerateID(),
		ActionId:     action.Id,
		GetHttpReq:   req.GetHttpReq,
		PostHttpReq:  req.PostHttpReq,
		PutHttpReq:   req.PutHttpReq,
		DelHttpReq:   req.DelHttpReq,
		ApiKey:       req.ApiKey,
		TextResponse: req.TextResponse,
	}

	err = a.actionHttpRepository.Create(ctx, actionHttp)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return actionHttp, nil
}

func (a *actionHttpService) FindActionHttpByID(ctx context.Context, actionId int64) (*model.ActionHttp, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":      ctx,
		"actionId": actionId,
	})

	if actionId <= 0 {
		return nil, constant.ErrInvalidArgument
	}

	actionHttp, err := a.actionHttpRepository.FindByActionID(ctx, actionId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if actionHttp == nil {
		return nil, constant.ErrNotFound
	}

	return actionHttp, nil
}

func (a *actionHttpService) UpdateActionHttp(ctx context.Context, actionId int64, req model.UpdateActionHttpRequest) (*model.ActionHttp, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":      ctx,
		"actionId": actionId,
		"req":      req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	actionHttp, err := a.FindActionHttpByID(ctx, actionId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	actionHttp.GetHttpReq = req.GetHttpReq
	actionHttp.PostHttpReq = req.PostHttpReq
	actionHttp.PutHttpReq = req.PutHttpReq
	actionHttp.DelHttpReq = req.DelHttpReq
	actionHttp.ApiKey = req.ApiKey
	actionHttp.TextResponse = req.TextResponse

	err = a.actionHttpRepository.Update(ctx, actionHttp)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return actionHttp, nil
}

func (a *actionHttpService) DeleteActionHttp(ctx context.Context, actionId int64) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":      ctx,
		"actionId": actionId,
	})

	_, err := a.FindActionHttpByID(ctx, actionId)
	if err != nil {
		log.Error(err)
		return false, err
	}

	err = a.actionHttpRepository.Delete(ctx, actionId)
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}
