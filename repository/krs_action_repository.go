package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type krsActionRepository struct {
	db *gorm.DB
}

func NewKrsActionRepository(db *gorm.DB) model.KrsActionRepository {
	return &krsActionRepository{
		db: db,
	}
}

func (k *krsActionRepository) Create(ctx context.Context, krsAction *model.KrsAction) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"krsAction": krsAction,
	})

	err := k.db.WithContext(ctx).Create(krsAction).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (k *krsActionRepository) FindAll(ctx context.Context) ([]*model.KrsAction, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	var krsActions []*model.KrsAction
	res := k.db.WithContext(ctx).Find(&krsActions)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return krsActions, nil
}

func (k *krsActionRepository) FindByID(ctx context.Context, id string) (*model.KrsAction, error) {
	if id == "" {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	krsAction := &model.KrsAction{}
	err := k.db.WithContext(ctx).Where("id = ?", id).Take(krsAction).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return krsAction, nil
}

func (k *krsActionRepository) Update(ctx context.Context, krsAction *model.KrsAction) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"krsAction": krsAction,
	})

	err := k.db.WithContext(ctx).Select(
		"name",
		"get_http_req",
		"api_key",
	).Updates(krsAction).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (k *krsActionRepository) Delete(ctx context.Context, id string) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := k.db.WithContext(ctx).Where("id = ?", id).Delete(&model.KrsAction{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
