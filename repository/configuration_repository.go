package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type configurationRepository struct {
	db *gorm.DB
}

func NewConfigurationRepository(db *gorm.DB) model.ConfigurationRepository {
	return &configurationRepository{
		db: db,
	}
}

func (c *configurationRepository) Create(ctx context.Context, configuration *model.Configuration) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":           ctx,
		"configuration": configuration,
	})

	err := c.db.WithContext(ctx).Create(configuration).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (c *configurationRepository) FindAll(ctx context.Context) ([]*model.Configuration, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	var configs []*model.Configuration

	res := c.db.WithContext(ctx).Find(&configs)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return configs, nil
}

func (c *configurationRepository) FindByID(ctx context.Context, id string) (*model.Configuration, error) {
	if id == "" {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	conf := &model.Configuration{}
	err := c.db.WithContext(ctx).Take(conf, "id = ?", id).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return conf, nil
}

func (c *configurationRepository) Update(ctx context.Context, id string, configuration *model.Configuration) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":           ctx,
		"id":            id,
		"configuration": configuration,
	})

	err := c.db.WithContext(ctx).Select(
		"diet_classifier_epoch",
		"fallback_classifier_treshold",
		"response_selector_epoch",
		"ted_policy_epoch",
		"fallback_utterance_id",
		"fallback_treshold",
		"unexpected_intent_policy_epoch",
	).Updates(configuration).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
