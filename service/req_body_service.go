package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"cbupnvj/repository"
	"context"

	"github.com/sirupsen/logrus"
)

type reqBodyService struct {
	reqBodyRepository    model.ReqBodyRepository
	actionHttpRepository model.ActionHttpRepository
	gormTransactioner    repository.GormTransactioner
}

func NewReqBodyService(reqBodyRepository model.ReqBodyRepository, actionHttpRepository model.ActionHttpRepository, gormTransactioner repository.GormTransactioner) model.ReqBodyService {
	return &reqBodyService{
		reqBodyRepository:    reqBodyRepository,
		actionHttpRepository: actionHttpRepository,
		gormTransactioner:    gormTransactioner,
	}
}

func (r *reqBodyService) CreateReqBody(ctx context.Context, req model.CreateReqBodyActionArrayRequest) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return false, constant.HttpValidationOrInternalErr(err)
	}

	actionHttpID, err := r.actionHttpRepository.FindByID(ctx, req.ActionHttpId)
	if err != nil {
		log.Error(err)
		return false, err
	}

	if actionHttpID == nil {
		log.Error(constant.ErrNotFound)
		return false, constant.ErrNotFound
	}

	if actionHttpID.PostHttpReq == "" || actionHttpID.PutHttpReq == "" {
		log.Error(constant.ErrFieldEmpty)
		return false, constant.ErrFieldEmpty
	}

	if len(req.PostFields) <= 0 && len(req.PutFields) <= 0 {
		log.Error("both fields are not more than 0")
		return false, constant.ErrInvalidArgument
	}

	tx := r.gormTransactioner.Begin(ctx)
	if len(req.PostFields) > 0 {
		for i := range req.PostFields {
			if err := req.PostFields[i].Validate(); err != nil {
				log.Error(err)
				r.gormTransactioner.Rollback(tx)
				return false, constant.HttpValidationOrInternalErr(err)
			}

			reqBody := &model.ReqBody{
				Id:           helper.GenerateID(),
				ActionHttpId: actionHttpID.Id,
				ReqName:      req.PostFields[i].ReqName,
				DataType:     req.PostFields[i].DataType,
				Method:       model.HttpMethodPost,
			}

			err = r.reqBodyRepository.Create(ctx, tx, reqBody)
			if err != nil {
				log.Error(err)
				r.gormTransactioner.Rollback(tx)
				return false, err
			}
		}
	}

	if len(req.PutFields) > 0 {
		for i := range req.PutFields {
			if err := req.PutFields[i].Validate(); err != nil {
				log.Error(err)
				r.gormTransactioner.Rollback(tx)
				return false, constant.HttpValidationOrInternalErr(err)
			}

			reqBody := &model.ReqBody{
				Id:           helper.GenerateID(),
				ActionHttpId: actionHttpID.Id,
				ReqName:      req.PutFields[i].ReqName,
				DataType:     req.PutFields[i].DataType,
				Method:       model.HttpMethodPut,
			}

			err = r.reqBodyRepository.Create(ctx, tx, reqBody)
			if err != nil {
				log.Error(err)
				r.gormTransactioner.Rollback(tx)
				return false, err
			}
		}
	}

	if err = r.gormTransactioner.Commit(tx); err != nil {
		log.Error(err)
		r.gormTransactioner.Rollback(tx)
		return false, err
	}

	return true, nil
}

func (r *reqBodyService) FindAllReqBodyByActionHttpID(ctx context.Context, actionHttpID string, method string) ([]*model.ReqBody, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":          ctx,
		"actionHttpID": actionHttpID,
		"method":       method,
	})

	var m model.HttpMethod
	if method == string(model.HttpMethodPost) {
		m = model.HttpMethodPost
	} else {
		m = model.HttpMethodPut
	}
	reqBodies, err := r.reqBodyRepository.FindAll(ctx, actionHttpID, m)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return reqBodies, nil
}

func (r *reqBodyService) FindByID(ctx context.Context, id string) (*model.ReqBody, error) {
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

func (r *reqBodyService) UpdateReqBody(ctx context.Context, id, actionHttpID string, req model.UpdateReqBodyRequest) (*model.ReqBody, error) {
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

func (r *reqBodyService) DeleteReqBody(ctx context.Context, id, actionHttpID string) (bool, error) {
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
