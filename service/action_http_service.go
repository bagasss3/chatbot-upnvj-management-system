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
	reqBodyRepository    model.ReqBodyRepository
}

func NewActionHttpService(actionHttpRepository model.ActionHttpRepository, reqBodyRepository model.ReqBodyRepository) model.ActionHttpService {
	return &actionHttpService{
		actionHttpRepository: actionHttpRepository,
		reqBodyRepository:    reqBodyRepository,
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

func (a *actionHttpService) FindAllActionHttp(ctx context.Context, name string) ([]*model.ActionHttp, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"name": name,
	})

	actionHttps, err := a.actionHttpRepository.FindAll(ctx, name)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return actionHttps, nil
}

func (a *actionHttpService) FindAllWithReqBodies(ctx context.Context) ([]*model.ActionHttp, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	actionHttps, err := a.actionHttpRepository.FindAllWithReqBodies(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return actionHttps, nil
}

func (a *actionHttpService) FindActionHttpByID(ctx context.Context, id string) (*model.ActionHttp, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	if id == "0" {
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

func (a *actionHttpService) UpdateActionHttp(ctx context.Context, id string, req model.CreateUpdateActionHttpRequest) (*model.ActionHttp, error) {
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

	if req.PostHttpReq == "" {
		actionPost, err := a.reqBodyRepository.FindAll(ctx, actionHttp.Id, model.HttpMethodPost)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		if len(actionPost) > 0 {
			log.Error("Post http can't be empty")
			return nil, constant.ErrFieldEmpty
		}
	}

	if req.PutHttpReq == "" {
		actionPut, err := a.reqBodyRepository.FindAll(ctx, actionHttp.Id, model.HttpMethodPut)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		if len(actionPut) > 0 {
			log.Error("Put http can't be empty")
			return nil, constant.ErrFieldEmpty
		}
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

func (a *actionHttpService) DeleteActionHttp(ctx context.Context, id string) (bool, error) {
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
