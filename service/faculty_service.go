package service

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type facultyService struct {
	facultyRepository model.FacultyRepository
}

func NewFacultyService(facultyRepository model.FacultyRepository) model.FacultyService {
	return &facultyService{
		facultyRepository: facultyRepository,
	}
}

func (f *facultyService) FindAllFaculty(ctx context.Context) ([]*model.Faculty, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	faculties, err := f.facultyRepository.FindAll(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return faculties, nil
}

func (f *facultyService) FindByIDFaculty(ctx context.Context, id string) (*model.Faculty, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	if id == "" {
		return nil, constant.ErrInvalidArgument
	}

	faculty, err := f.facultyRepository.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if faculty == nil {
		return nil, constant.ErrNotFound
	}

	return faculty, nil
}
