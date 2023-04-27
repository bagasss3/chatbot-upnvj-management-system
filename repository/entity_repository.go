package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type entityRepository struct {
	db *gorm.DB
}

func NewEntityRepository(db *gorm.DB) model.EntityRepository {
	return &entityRepository{
		db: db,
	}
}

func (e *entityRepository) Create(ctx context.Context, entity *model.Entity) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":    ctx,
		"entity": entity,
	})

	err := e.db.WithContext(ctx).Create(entity).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (e *entityRepository) FindAll(ctx context.Context, intentId string) ([]*model.Entity, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":      ctx,
		"intentId": intentId,
	})

	var entities []*model.Entity
	res := e.db.WithContext(ctx).Where("intent_id = ?", intentId).Find(&entities)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return entities, nil
}

func (e *entityRepository) FindByID(ctx context.Context, id string) (*model.Entity, error) {
	if id == "" {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	entity := &model.Entity{}
	err := e.db.WithContext(ctx).Where("id = ?", id).Take(entity).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return entity, nil
}

func (e *entityRepository) Delete(ctx context.Context, id string) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := e.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Entity{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
