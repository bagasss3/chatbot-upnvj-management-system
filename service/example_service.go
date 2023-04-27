package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type exampleService struct {
	exampleRepository model.ExampleRepository
	intentRepository  model.IntentRepository
}

func NewExampleService(exampleRepository model.ExampleRepository, intentRepository model.IntentRepository) model.ExampleService {
	return &exampleService{
		exampleRepository: exampleRepository,
		intentRepository:  intentRepository,
	}
}

func (e *exampleService) CreateExample(ctx context.Context, req model.CreateExampleRequest) (*model.Example, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	intent, err := e.intentRepository.FindByID(ctx, req.IntentID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if intent == nil {
		log.Error(constant.ErrNotFound)
		return nil, constant.ErrNotFound
	}

	example := &model.Example{
		Id:       helper.GenerateID(),
		Example:  req.Example,
		IntentId: intent.Id,
	}

	err = e.exampleRepository.Create(ctx, example)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return example, nil
}

func (e *exampleService) FindAllExampleByIntentID(ctx context.Context, id string) ([]*model.Example, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	if id == "" {
		log.Error(constant.ErrNotFound)
		return nil, constant.ErrNotFound
	}

	examples, err := e.exampleRepository.FindAllByIntentID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return examples, nil
}

func (e *exampleService) FindExampleByIntentID(ctx context.Context, intentId, exampleId string) (*model.Example, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"IntentId":  intentId,
		"ExampleId": exampleId,
	})

	example, err := e.exampleRepository.FindByID(ctx, intentId, exampleId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if example == nil {
		log.Error(constant.ErrNotFound)
		return nil, constant.ErrNotFound
	}

	return example, nil
}

func (e *exampleService) UpdateExample(ctx context.Context, intentId, exampleId string, req model.UpdateExampleRequest) (*model.Example, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"IntentId":  intentId,
		"ExampleId": exampleId,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	example, err := e.FindExampleByIntentID(ctx, intentId, exampleId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	example.Example = req.Example
	err = e.exampleRepository.Update(ctx, exampleId, example)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return example, nil
}

func (e *exampleService) DeleteExample(ctx context.Context, intentId, exampleId string) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"IntentId":  intentId,
		"ExampleId": exampleId,
	})

	_, err := e.FindExampleByIntentID(ctx, intentId, exampleId)
	if err != nil {
		log.Error(err)
		return false, err
	}

	err = e.exampleRepository.Delete(ctx, exampleId)
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}
