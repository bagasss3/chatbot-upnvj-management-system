package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type utteranceRepository struct {
	db *gorm.DB
}

func NewUtteranceRepository(db *gorm.DB) model.UtteranceRepository {
	return &utteranceRepository{
		db: db,
	}
}

func (u *utteranceRepository) Create(ctx context.Context, utterance *model.Utterance) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"utterance": utterance,
	})

	err := u.db.WithContext(ctx).Create(utterance).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (u *utteranceRepository) FindByID(ctx context.Context, id string) (*model.Utterance, error) {
	if id == "" {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	utterance := &model.Utterance{}
	err := u.db.WithContext(ctx).Take(utterance, "id = ?", id).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return utterance, nil
}

func (u *utteranceRepository) FindAll(ctx context.Context, name string) ([]*model.Utterance, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"name": name,
	})

	var utterances []*model.Utterance
	var res *gorm.DB
	if name == "" {
		res = u.db.WithContext(ctx).Find(&utterances)
	} else {
		res = u.db.WithContext(ctx).Where("name LIKE ?", "%"+name+"%").Find(&utterances)
	}
	if res.Error != nil {
		log.Error(res.Error)
		return utterances, res.Error
	}

	return utterances, nil
}

func (u *utteranceRepository) Update(ctx context.Context, id string, utterance *model.Utterance) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"id":        id,
		"utterance": utterance,
	})

	err := u.db.WithContext(ctx).Select(
		"name",
		"response",
	).Updates(utterance).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (u *utteranceRepository) Delete(ctx context.Context, id string) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := u.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Utterance{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
