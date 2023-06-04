package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type trainingHistoryService struct {
	trainingHistoryRepository model.TrainingHistoryRepository
	userRepository            model.UserRepository
}

func NewTrainingHistoryService(trainingHistoryRepository model.TrainingHistoryRepository, userRepository model.UserRepository) model.TrainingHistoryService {
	return &trainingHistoryService{
		trainingHistoryRepository: trainingHistoryRepository,
		userRepository:            userRepository,
	}
}

func (th *trainingHistoryService) CreateTrainingHistory(ctx context.Context, req model.CreateTrainingHistoryRequest) (*model.TrainingHistory, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	user, err := th.userRepository.FindByID(ctx, req.UserId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if user == nil {
		log.Error("user not found")
		return nil, constant.ErrNotFound
	}

	if req.Status != model.StatusDone && req.Status != model.StatusFailed {
		log.Error("invalid status training")
		return nil, constant.ErrInvalidArgument
	}

	trainingHistory := &model.TrainingHistory{
		Id:        helper.GenerateID(),
		UserId:    req.UserId,
		TotalTime: req.TotalTime,
		Status:    req.Status,
		User:      user,
	}

	err = th.trainingHistoryRepository.Create(ctx, trainingHistory)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return trainingHistory, nil
}

func (th *trainingHistoryService) FindAllTrainingHistory(ctx context.Context) ([]*model.TrainingHistory, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	trainingHistories, err := th.trainingHistoryRepository.FindAll(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return trainingHistories, nil
}
