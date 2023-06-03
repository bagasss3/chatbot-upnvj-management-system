package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type facultyRepository struct {
	db *gorm.DB
}

func NewFacultyRepository(db *gorm.DB) model.FacultyRepository {
	return &facultyRepository{
		db: db,
	}
}

func (f *facultyRepository) FindAll(ctx context.Context) ([]*model.Faculty, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	var Faculties []*model.Faculty
	res := f.db.WithContext(ctx).Preload("Majors").Find(&Faculties)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return Faculties, nil
}

func (f *facultyRepository) FindByID(ctx context.Context, id string) (*model.Faculty, error) {
	if id == "" {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	faculty := &model.Faculty{}
	err := f.db.WithContext(ctx).Take(faculty, "id = ?", id).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return faculty, nil
}
