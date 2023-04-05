package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type exampleRepository struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) model.ExampleRepository {
	return &exampleRepository{
		db: db,
	}
}

func (e *exampleRepository) Create(ctx context.Context, example *model.Example) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"example": example,
	})

	err := e.db.WithContext(ctx).Create(example).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (e *exampleRepository) FindByID(ctx context.Context, intentId, exampleId int64) (*model.Example, error) {
	if exampleId <= 0 || intentId <= 0 {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"intentId":  intentId,
		"exampleId": exampleId,
	})

	example := &model.Example{}
	err := e.db.WithContext(ctx).Where("id = ?", exampleId).Where("intent_id = ?", intentId).Take(example).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return example, nil
}

func (e *exampleRepository) FindAllByIntentID(ctx context.Context, intentId int64) ([]*model.Example, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":      ctx,
		"intentId": intentId,
	})

	var examples []*model.Example
	res := e.db.WithContext(ctx).Where("intent_id = ?", intentId).Find(&examples)
	if res.Error != nil {
		log.Error(res.Error)
		return examples, res.Error
	}

	return examples, nil
}

func (e *exampleRepository) Update(ctx context.Context, id int64, example *model.Example) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"id":      id,
		"example": example,
	})

	err := e.db.WithContext(ctx).Select(
		"example",
	).Updates(example).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (e *exampleRepository) Delete(ctx context.Context, id int64) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := e.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Example{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
