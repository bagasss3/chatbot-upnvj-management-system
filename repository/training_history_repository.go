package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type trainingHistoryRepository struct {
	db *gorm.DB
}

func NewTrainingHistoryRepository(db *gorm.DB) model.TrainingHistoryRepository {
	return &trainingHistoryRepository{
		db: db,
	}
}

func (th *trainingHistoryRepository) Create(ctx context.Context, trainingHistory *model.TrainingHistory) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":             ctx,
		"trainingHistory": trainingHistory,
	})

	err := th.db.WithContext(ctx).Create(trainingHistory).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (th *trainingHistoryRepository) FindAll(ctx context.Context) ([]*model.TrainingHistory, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	var trainingHistories []*model.TrainingHistory
	res := th.db.WithContext(ctx).Preload("User").Order("created_at DESC").Find(&trainingHistories)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return trainingHistories, nil
}
