package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type intentRepository struct {
	db *gorm.DB
}

func NewIntentRepository(db *gorm.DB) model.IntentRepository {
	return &intentRepository{
		db: db,
	}
}

func (i *intentRepository) Create(ctx context.Context, intent *model.Intent) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":    ctx,
		"intent": intent,
	})

	err := i.db.WithContext(ctx).Create(intent).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (i *intentRepository) FindByID(ctx context.Context, id int64) (*model.Intent, error) {
	if id <= 0 {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	intent := &model.Intent{}
	err := i.db.WithContext(ctx).Take(intent, "id = ?", id).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return intent, nil
}

func (i *intentRepository) FindAll(ctx context.Context) ([]*model.Intent, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	var intents []*model.Intent
	res := i.db.WithContext(ctx).Find(&intents)
	if res.Error != nil {
		log.Error(res.Error)
		return intents, res.Error
	}

	return intents, nil
}

func (i *intentRepository) Update(ctx context.Context, id int64, intent *model.Intent) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":    ctx,
		"id":     id,
		"intent": intent,
	})

	err := i.db.WithContext(ctx).Select(
		"name",
	).Updates(intent).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (i *intentRepository) Delete(ctx context.Context, id int64) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := i.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Intent{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
