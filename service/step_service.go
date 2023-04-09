package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type stepService struct {
	storyRepository      model.StoryRepository
	stepRepository       model.StepRepository
	intentRepository     model.IntentRepository
	utteranceRepository  model.UtteranceRepository
	actionHttpRepository model.ActionHttpRepository
}

func NewStepService(storyRepository model.StoryRepository, stepRepository model.StepRepository,
	intentRepository model.IntentRepository, utteranceRepository model.UtteranceRepository, actionHttpRepository model.ActionHttpRepository) model.StepService {
	return &stepService{
		storyRepository:      storyRepository,
		stepRepository:       stepRepository,
		intentRepository:     intentRepository,
		utteranceRepository:  utteranceRepository,
		actionHttpRepository: actionHttpRepository,
	}
}

func (s *stepService) CreateStep(ctx context.Context, req model.CreateStepRequest) (*model.Step, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	story, err := s.storyRepository.FindByID(ctx, req.StoryId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if story == nil {
		log.Error("story not found")
		return nil, constant.ErrNotFound
	}

	err = s.ValidateStep(ctx, req.ResponseId, req.Type)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	step := &model.Step{
		Id:         helper.GenerateID(),
		StoryId:    req.StoryId,
		ResponseId: req.ResponseId,
		Type:       req.Type,
	}

	err = s.stepRepository.Create(ctx, step)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return step, nil
}

func (s *stepService) FindAllStepByStoryID(ctx context.Context, storyId int64) ([]*model.Step, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"storyId": storyId,
	})

	steps, err := s.stepRepository.FindAll(ctx, storyId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return steps, nil
}

func (s *stepService) FindStepByID(ctx context.Context, id, storyId int64) (*model.Step, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"id":      id,
		"storyId": storyId,
	})

	if id <= 0 || storyId <= 0 {
		return nil, constant.ErrInvalidArgument
	}

	step, err := s.stepRepository.FindByID(ctx, id, storyId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if step == nil {
		return nil, constant.ErrNotFound
	}

	return step, nil
}

func (s *stepService) UpdateStep(ctx context.Context, id, storyId int64, req model.UpdateStepRequest) (*model.Step, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"id":      id,
		"storyId": storyId,
		"req":     req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	err := s.ValidateStep(ctx, req.ResponseId, req.Type)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	step, err := s.FindStepByID(ctx, id, storyId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	step.ResponseId = req.ResponseId
	step.Type = req.Type

	err = s.stepRepository.Update(ctx, id, step)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return step, nil
}

func (s *stepService) DeleteStep(ctx context.Context, id, storyId int64) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	_, err := s.FindStepByID(ctx, id, storyId)
	if err != nil {
		log.Error(err)
		return false, err
	}

	err = s.stepRepository.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}

func (s *stepService) ValidateStep(ctx context.Context, responseId int64, stepType model.StepType) error {
	if stepType != model.StepAction && stepType != model.StepIntent && stepType != model.StepUtterance {
		return constant.ErrInvalidArgument
	}

	switch stepType {
	case model.StepAction:
		action, err := s.actionHttpRepository.FindByID(ctx, responseId)
		if err != nil {
			return err
		}

		if action == nil {
			return constant.ErrNotFound
		}
	case model.StepIntent:
		intent, err := s.intentRepository.FindByID(ctx, responseId)
		if err != nil {
			return err
		}

		if intent == nil {
			return constant.ErrNotFound
		}
	case model.StepUtterance:
		utterance, err := s.utteranceRepository.FindByID(ctx, responseId)
		if err != nil {
			return err
		}

		if utterance == nil {
			return constant.ErrNotFound
		}
	}

	return nil
}
