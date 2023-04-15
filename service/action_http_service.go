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
}

func NewActionHttpService(actionHttpRepository model.ActionHttpRepository) model.ActionHttpService {
	return &actionHttpService{
		actionHttpRepository: actionHttpRepository,
	}
}

func (a *actionHttpService) CreateActionHttp(ctx context.Context, req model.CreateUpdateActionHttpRequest) (*model.ActionHttp, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	actionHttp := &model.ActionHttp{
		Id:           helper.GenerateID(),
		Name:         req.Name,
		GetHttpReq:   req.GetHttpReq,
		PostHttpReq:  req.PostHttpReq,
		PutHttpReq:   req.PutHttpReq,
		DelHttpReq:   req.DelHttpReq,
		ApiKey:       req.ApiKey,
		TextResponse: req.TextResponse,
	}

	err := a.actionHttpRepository.Create(ctx, actionHttp)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return actionHttp, nil
}

func (a *actionHttpService) FindAllActionHttp(ctx context.Context) ([]*model.ActionHttp, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	actionHttps, err := a.actionHttpRepository.FindAll(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return actionHttps, nil
}

func (a *actionHttpService) FindActionHttpByID(ctx context.Context, id int64) (*model.ActionHttp, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	if id <= 0 {
		return nil, constant.ErrInvalidArgument
	}

	actionHttp, err := a.actionHttpRepository.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if actionHttp == nil {
		return nil, constant.ErrNotFound
	}

	return actionHttp, nil
}

func (a *actionHttpService) UpdateActionHttp(ctx context.Context, id int64, req model.CreateUpdateActionHttpRequest) (*model.ActionHttp, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
		"req": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	actionHttp, err := a.FindActionHttpByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	actionHttp.Name = req.Name
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

func (a *actionHttpService) DeleteActionHttp(ctx context.Context, id int64) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	_, err := a.FindActionHttpByID(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	err = a.actionHttpRepository.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}

func (a *actionHttpService) CountAllActionHttp(ctx context.Context) (int64, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	count, err := a.actionHttpRepository.CountAll(ctx)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return count, nil
}
