package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type utteranceService struct {
	utteranceRepository model.UtteranceRepository
}

func NewUtteranceService(utteranceRepository model.UtteranceRepository) model.UtteranceService {
	return &utteranceService{
		utteranceRepository: utteranceRepository,
	}
}

func (u *utteranceService) CreateUtterance(ctx context.Context, req model.CreateUpdateUtteranceRequest) (*model.Utterance, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	utterance := &model.Utterance{
		Id:       helper.GenerateID(),
		Name:     req.Name,
		Response: req.Response,
	}

	err := u.utteranceRepository.Create(ctx, utterance)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return utterance, nil
}

func (u *utteranceService) FindAllUtterance(ctx context.Context, name string) ([]*model.Utterance, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"name": name,
	})

	utterances, err := u.utteranceRepository.FindAll(ctx, name)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return utterances, nil
}

func (u *utteranceService) FindUtteranceByID(ctx context.Context, id string) (*model.Utterance, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	if id == "" {
		return nil, constant.ErrInvalidArgument
	}

	utterance, err := u.utteranceRepository.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if utterance == nil {
		return nil, constant.ErrNotFound
	}

	return utterance, nil
}

func (u *utteranceService) UpdateUtterance(ctx context.Context, id string, req model.CreateUpdateUtteranceRequest) (*model.Utterance, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
		"req": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	utterance, err := u.FindUtteranceByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	utterance.Name = req.Name
	utterance.Response = req.Response
	err = u.utteranceRepository.Update(ctx, utterance.Id, utterance)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return utterance, nil
}

func (u *utteranceService) DeleteUtterance(ctx context.Context, id string) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	_, err := u.FindUtteranceByID(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	err = u.utteranceRepository.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}

func (u *utteranceService) CountAllUtterance(ctx context.Context) (int64, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	count, err := u.utteranceRepository.CountAll(ctx)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return count, nil
}
