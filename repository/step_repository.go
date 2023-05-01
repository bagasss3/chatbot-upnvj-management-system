package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type stepRepository struct {
	db *gorm.DB
}

func NewStepRepository(db *gorm.DB) model.StepRepository {
	return &stepRepository{
		db: db,
	}
}

func (s *stepRepository) Create(ctx context.Context, step *model.Step) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"step": step,
	})

	err := s.db.WithContext(ctx).Create(step).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *stepRepository) FindAll(ctx context.Context, storyId string) ([]*model.Step, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	var steps []*model.Step
	res := s.db.WithContext(ctx).Where("story_id = ?", storyId).Find(&steps)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return steps, nil
}

func (s *stepRepository) FindByID(ctx context.Context, id, storyId string) (*model.Step, error) {
	if id == "" || storyId == "" {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"id":      id,
		"storyId": storyId,
	})

	step := &model.Step{}
	err := s.db.WithContext(ctx).Where("id = ?", id).Where("story_id = ?", storyId).Take(step).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return step, nil
}

func (s *stepRepository) Update(ctx context.Context, id string, step *model.Step) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"id":   id,
		"step": step,
	})

	err := s.db.WithContext(ctx).Select(
		"response_id",
		"type",
	).Updates(step).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *stepRepository) Delete(ctx context.Context, id string) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := s.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Step{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *stepRepository) DeleteAllByStoryID(ctx context.Context, tx *gorm.DB, storyId string) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"storyId": storyId,
	})

	err := tx.WithContext(ctx).Where("story_id = ?", storyId).Delete(&model.Step{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
