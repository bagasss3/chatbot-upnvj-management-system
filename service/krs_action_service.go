package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type krsActionService struct {
	krsActionRepository model.KrsActionRepository
}

func NewKrsActionService(krsActionRepository model.KrsActionRepository) model.KrsActionService {
	return &krsActionService{
		krsActionRepository: krsActionRepository,
	}
}

func (k *krsActionService) CreateKrsAction(ctx context.Context, req model.CreateUpdateKrsActionRequest) (*model.KrsAction, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	krsAction := &model.KrsAction{
		Id:         helper.GenerateID(),
		Name:       req.Name,
		GetHttpReq: req.GetHttpReq,
		ApiKey:     req.ApiKey,
	}

	err := k.krsActionRepository.Create(ctx, krsAction)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return krsAction, nil
}

func (k *krsActionService) FindAllKrsAction(ctx context.Context) ([]*model.KrsAction, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	krsActions, err := k.krsActionRepository.FindAll(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return krsActions, nil
}

func (k *krsActionService) FindKrsActionByID(ctx context.Context, id int64) (*model.KrsAction, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	if id <= 0 {
		return nil, constant.ErrInvalidArgument
	}

	krsAction, err := k.krsActionRepository.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if krsAction == nil {
		return nil, constant.ErrNotFound
	}

	return krsAction, nil
}

func (k *krsActionService) UpdateKrsAction(ctx context.Context, id int64, req model.CreateUpdateKrsActionRequest) (*model.KrsAction, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
		"req": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	krsAction, err := k.krsActionRepository.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	krsAction.Name = req.Name
	krsAction.GetHttpReq = req.GetHttpReq
	krsAction.ApiKey = req.ApiKey

	err = k.krsActionRepository.Update(ctx, krsAction)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return krsAction, nil
}

func (k *krsActionService) DeleteKrsAction(ctx context.Context, id int64) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	_, err := k.FindKrsActionByID(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	err = k.krsActionRepository.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}
