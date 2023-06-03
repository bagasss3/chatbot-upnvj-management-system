package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type majorRepository struct {
	db *gorm.DB
}

func NewMajorRepository(db *gorm.DB) model.MajorRepository {
	return &majorRepository{
		db: db,
	}
}

func (m *majorRepository) FindAll(ctx context.Context) ([]*model.Major, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	var majors []*model.Major
	res := m.db.WithContext(ctx).Find(&majors)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return majors, nil
}

func (m *majorRepository) FindByID(ctx context.Context, id string) (*model.Major, error) {
	if id == "" {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	major := &model.Major{}
	err := m.db.WithContext(ctx).Take(major, "id = ?", id).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return major, nil
}
