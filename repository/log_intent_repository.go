package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type logIntentRepository struct {
	db *gorm.DB
}

func NewLogIntentRepository(db *gorm.DB) model.LogIntentRepository {
	return &logIntentRepository{
		db: db,
	}
}

func (li *logIntentRepository) Create(ctx context.Context, logIntent *model.LogIntent) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"logIntent": logIntent,
	})

	err := li.db.WithContext(ctx).Create(logIntent).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (li *logIntentRepository) FindAll(ctx context.Context) ([]*model.LogIntent, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	var logintents []*model.LogIntent
	res := li.db.WithContext(ctx).Find(&logintents)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return logintents, nil
}

func (li *logIntentRepository) FindByIntentID(ctx context.Context, intentId string) (*model.LogIntent, error) {
	if intentId == "" {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx":      ctx,
		"intentId": intentId,
	})

	logIntent := &model.LogIntent{}
	err := li.db.WithContext(ctx).Where("intent_id = ?", intentId).Take(logIntent).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return logIntent, nil
}

func (li *logIntentRepository) Update(ctx context.Context, intentId string, logIntent *model.LogIntent) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"intentId":  intentId,
		"logIntent": logIntent,
	})

	err := li.db.WithContext(ctx).Select(
		"mention",
	).Updates(logIntent).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
