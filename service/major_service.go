package service

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type majorService struct {
	majorRepository model.MajorRepository
}

func NewMajorService(majorRepository model.MajorRepository) model.MajorService {
	return &majorService{
		majorRepository: majorRepository,
	}
}

func (m *majorService) FindAllMajor(ctx context.Context) ([]*model.Major, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	majors, err := m.majorRepository.FindAll(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return majors, nil
}

func (m *majorService) FindByIDMajor(ctx context.Context, id string) (*model.Major, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	if id == "" {
		return nil, constant.ErrInvalidArgument
	}

	major, err := m.majorRepository.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if major == nil {
		return nil, constant.ErrNotFound
	}

	return major, nil
}
