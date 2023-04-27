package repository

import (
	"cbupnvj/model"
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) model.SessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (s *sessionRepository) Create(ctx context.Context, session *model.Session) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"session": session,
	})

	err := s.db.WithContext(ctx).Create(session).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *sessionRepository) FindByID(ctx context.Context, id string) (*model.Session, error) {
	if id == "" {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	session := &model.Session{}
	err := s.db.WithContext(ctx).Take(session, "id = ?", id).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return session, nil
}

func (s *sessionRepository) FindByRefreshToken(ctx context.Context, refreshToken string) (*model.Session, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":         ctx,
		"accessToken": refreshToken,
	})

	id := string("")
	err := s.db.Model(model.Session{}).Select("id").Take(&id, "refresh_token = ?", refreshToken).Error
	switch err {
	case nil:
		return s.FindByID(ctx, id)
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}
}

func (s *sessionRepository) RefreshToken(ctx context.Context, session *model.Session) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"session": session,
	})

	session.UpdatedAt = time.Now()
	err := s.db.WithContext(ctx).Model(model.Session{}).Select(
		"access_token",
		"access_token_expired_at",
		"refresh_token",
		"refresh_token_expired_at",
		"updated_at",
	).Where("id = ?", session.Id).Updates(session).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
