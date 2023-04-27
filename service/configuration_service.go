package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type configurationService struct {
	configurationRepository model.ConfigurationRepository
	utteranceRepository     model.UtteranceRepository
}

func NewConfigurationService(configurationRepository model.ConfigurationRepository, utteranceRepository model.UtteranceRepository) model.ConfigurationService {
	return &configurationService{
		configurationRepository: configurationRepository,
		utteranceRepository:     utteranceRepository,
	}
}

func (c *configurationService) CreateConfiguration(ctx context.Context, req model.CreateUpdateConfigurationRequest) (*model.Configuration, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	utterance, err := c.utteranceRepository.FindByID(ctx, req.FallbackUtteranceId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if utterance == nil {
		log.Error(constant.ErrNotFound)
		return nil, constant.ErrNotFound
	}

	conf := &model.Configuration{
		Id:                         helper.GenerateID(),
		DietClassifierEpoch:        req.DietClassifierEpoch,
		FallbackClassifierTreshold: req.FallbackClassifierTreshold,
		ResponseSelectorEpoch:      req.ResponseSelectorEpoch,
		TedPolicyEpoch:             req.TedPolicyEpoch,
		FallbackUtteranceId:        req.FallbackUtteranceId,
		FallbackTreshold:           req.FallbackTreshold,
	}

	err = c.configurationRepository.Create(ctx, conf)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return conf, nil
}

func (c *configurationService) FindConfiguration(ctx context.Context, id string) (*model.Configuration, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	if id == "" {
		return nil, constant.ErrInvalidArgument
	}

	conf, err := c.configurationRepository.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if conf == nil {
		log.Error(constant.ErrNotFound)
		return nil, constant.ErrNotFound
	}

	return conf, nil
}

func (c *configurationService) UpdateConfiguration(ctx context.Context, id string, req model.CreateUpdateConfigurationRequest) (*model.Configuration, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
		"req": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	conf, err := c.FindConfiguration(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	utterance, err := c.utteranceRepository.FindByID(ctx, req.FallbackUtteranceId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if utterance == nil {
		log.Error(constant.ErrNotFound)
		return nil, constant.ErrNotFound
	}

	conf.FallbackUtteranceId = req.FallbackUtteranceId
	conf.DietClassifierEpoch = req.DietClassifierEpoch
	conf.FallbackClassifierTreshold = req.FallbackClassifierTreshold
	conf.FallbackTreshold = req.FallbackTreshold
	conf.ResponseSelectorEpoch = req.ResponseSelectorEpoch
	conf.TedPolicyEpoch = req.TedPolicyEpoch

	err = c.configurationRepository.Update(ctx, conf.Id, conf)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return conf, nil
}
