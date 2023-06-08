package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type storyRepository struct {
	db *gorm.DB
}

func NewStoryRepository(db *gorm.DB) model.StoryRepository {
	return &storyRepository{
		db: db,
	}
}

func (s *storyRepository) Create(ctx context.Context, story *model.Story) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":   ctx,
		"story": story,
	})

	err := s.db.WithContext(ctx).Create(story).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *storyRepository) FindAll(ctx context.Context, name string) ([]*model.Story, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"name": name,
	})

	var stories []*model.Story
	var res *gorm.DB
	if name == "" {
		res = s.db.WithContext(ctx).Order("created_at DESC").Find(&stories)
	} else {
		res = s.db.WithContext(ctx).Where("story_title LIKE ?", "%"+name+"%").Order("created_at DESC").Find(&stories)
	}
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return stories, nil
}

func (s *storyRepository) FindByID(ctx context.Context, id string) (*model.Story, error) {
	if id == "" {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	story := &model.Story{}
	err := s.db.WithContext(ctx).Take(story, "id = ?", id).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return story, nil
}

func (s *storyRepository) Update(ctx context.Context, id string, story *model.Story) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":   ctx,
		"id":    id,
		"story": story,
	})

	err := s.db.WithContext(ctx).Select(
		"story_title",
	).Updates(story).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *storyRepository) Delete(ctx context.Context, id string) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := s.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Story{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *storyRepository) DeleteWithContext(ctx context.Context, tx *gorm.DB, id string) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := tx.WithContext(ctx).Where("id = ?", id).Delete(&model.Story{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *storyRepository) CountAll(ctx context.Context) (int64, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})
	var count int64
	err := s.db.WithContext(ctx).Model(&model.Story{}).Count(&count).Error
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return count, nil
}
