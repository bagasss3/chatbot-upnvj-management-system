package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type actionRepository struct {
	db *gorm.DB
}

func NewActionRepository(db *gorm.DB) model.ActionRepository {
	return &actionRepository{
		db: db,
	}
}

func (a *actionRepository) Create(ctx context.Context, action *model.Action) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":    ctx,
		"action": action,
	})

	err := a.db.WithContext(ctx).Create(action).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (a *actionRepository) FindByID(ctx context.Context, id int64) (*model.Action, error) {
	if id <= 0 {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	action := &model.Action{}
	err := a.db.WithContext(ctx).Take(action, "id = ?", id).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return action, nil
}

func (a *actionRepository) FindAll(ctx context.Context) ([]*model.Action, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	var actions []*model.Action
	res := a.db.WithContext(ctx).Find(&actions)
	if res.Error != nil {
		log.Error(res.Error)
		return actions, res.Error
	}

	return actions, nil
}

func (a *actionRepository) Update(ctx context.Context, id int64, action *model.Action) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":    ctx,
		"id":     id,
		"action": action,
	})

	err := a.db.WithContext(ctx).Select(
		"name",
	).Updates(action).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (a *actionRepository) Delete(ctx context.Context, id int64) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := a.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Action{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
