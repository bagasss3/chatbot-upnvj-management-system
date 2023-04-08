package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type reqBodyService struct {
	reqBodyRepository    model.ReqBodyRepository
	actionHttpRepository model.ActionHttpRepository
}

func NewReqBodyService(reqBodyRepository model.ReqBodyRepository, actionHttpRepository model.ActionHttpRepository) model.ReqBodyService {
	return &reqBodyService{
		reqBodyRepository:    reqBodyRepository,
		actionHttpRepository: actionHttpRepository,
	}
}

func (r *reqBodyService) CreateReqBody(ctx context.Context, req model.CreateReqBodyRequest) (*model.ReqBody, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	actionHttpID, err := r.actionHttpRepository.FindByID(ctx, req.ActionHttpId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if actionHttpID == nil {
		log.Error(constant.ErrNotFound)
		return nil, constant.ErrNotFound
	}

	reqBody := &model.ReqBody{
		Id:           helper.GenerateID(),
		ActionHttpId: actionHttpID.Id,
		ReqName:      req.ReqName,
	}

	err = r.reqBodyRepository.Create(ctx, reqBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return reqBody, nil
}

func (r *reqBodyService) FindAllReqBodyByActionHttpID(ctx context.Context, actionHttpID int64) ([]*model.ReqBody, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":          ctx,
		"actionHttpID": actionHttpID,
	})

	reqBodies, err := r.reqBodyRepository.FindAll(ctx, actionHttpID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return reqBodies, nil
}

func (r *reqBodyService) FindByID(ctx context.Context, id int64) (*model.ReqBody, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	reqBody, err := r.reqBodyRepository.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if reqBody == nil {
		return nil, constant.ErrNotFound
	}

	return reqBody, nil
}

func (r *reqBodyService) UpdateReqBody(ctx context.Context, id, actionHttpID int64, req model.UpdateReqBodyRequest) (*model.ReqBody, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":          ctx,
		"id":           id,
		"actionHttpID": actionHttpID,
		"req":          req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	actionHttp, err := r.actionHttpRepository.FindByID(ctx, actionHttpID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if actionHttp == nil {
		log.Error(constant.ErrNotFound)
		return nil, constant.ErrNotFound
	}

	reqBody, err := r.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if reqBody.ActionHttpId != actionHttp.Id {
		log.Error(err)
		return nil, constant.ErrInvalidArgument
	}

	reqBody.ReqName = req.ReqName
	err = r.reqBodyRepository.Update(ctx, id, reqBody)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return reqBody, nil
}

func (r *reqBodyService) DeleteReqBody(ctx context.Context, id, actionHttpID int64) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":          ctx,
		"id":           id,
		"actionHttpID": actionHttpID,
	})

	actionHttp, err := r.actionHttpRepository.FindByID(ctx, actionHttpID)
	if err != nil {
		log.Error(err)
		return false, err
	}

	if actionHttp == nil {
		log.Error(constant.ErrNotFound)
		return false, constant.ErrNotFound
	}

	reqBody, err := r.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	if reqBody.ActionHttpId != actionHttp.Id {
		log.Error(err)
		return false, constant.ErrInvalidArgument
	}

	err = r.reqBodyRepository.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}
