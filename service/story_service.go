package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"cbupnvj/repository"
	"context"

	"github.com/sirupsen/logrus"
)

type storyService struct {
	storyRepository   model.StoryRepository
	stepRepository    model.StepRepository
	gormTransactioner repository.GormTransactioner
}

func NewStoryService(storyRepository model.StoryRepository, stepRepository model.StepRepository, gormTransactioner repository.GormTransactioner) model.StoryService {
	return &storyService{
		storyRepository:   storyRepository,
		stepRepository:    stepRepository,
		gormTransactioner: gormTransactioner,
	}
}

func (s *storyService) CreateStory(ctx context.Context, req model.CreateUpdateStoryRequest) (*model.Story, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	story := &model.Story{
		Id:         helper.GenerateID(),
		StoryTitle: req.StoryTitle,
	}

	err := s.storyRepository.Create(ctx, story)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return story, nil
}

func (s *storyService) FindAllStory(ctx context.Context) ([]*model.Story, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	stories, err := s.storyRepository.FindAll(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return stories, nil
}

func (s *storyService) FindStoryByID(ctx context.Context, id string) (*model.Story, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	if id == "" {
		return nil, constant.ErrInvalidArgument
	}

	story, err := s.storyRepository.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if story == nil {
		return nil, constant.ErrNotFound
	}

	return story, nil
}

func (s *storyService) UpdateStory(ctx context.Context, id string, req model.CreateUpdateStoryRequest) (*model.Story, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
		"req": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	story, err := s.FindStoryByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	story.StoryTitle = req.StoryTitle

	err = s.storyRepository.Update(ctx, id, story)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return story, nil
}

func (s *storyService) DeleteStory(ctx context.Context, id string) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	_, err := s.FindStoryByID(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	tx := s.gormTransactioner.Begin(ctx)
	err = s.storyRepository.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		s.gormTransactioner.Rollback(tx)
		return false, err
	}

	err = s.stepRepository.DeleteAllByStoryID(ctx, id)
	if err != nil {
		log.Error(err)
		s.gormTransactioner.Rollback(tx)
		return false, err
	}

	if err = s.gormTransactioner.Commit(tx); err != nil {
		log.Error(err)
		s.gormTransactioner.Rollback(tx)
		return false, err
	}
	return true, nil
}
