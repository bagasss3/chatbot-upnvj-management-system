package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type fallbackChatLogRepository struct {
	db *gorm.DB
}

func NewFallbackChatLogRepository(db *gorm.DB) model.FallbackChatLogRepository {
	return &fallbackChatLogRepository{
		db: db,
	}
}

func (fcl *fallbackChatLogRepository) Create(ctx context.Context, fallback *model.FallbackChatLog) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":             ctx,
		"fallbackChatLog": fallback,
	})

	err := fcl.db.WithContext(ctx).Create(fallback).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (fcl *fallbackChatLogRepository) FindAll(ctx context.Context, page string) ([]*model.FallbackChatLog, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"page": page,
	})

	var fallbackChatLog []*model.FallbackChatLog
	var res *gorm.DB

	if page == string(model.FallbackPageDashboard) {
		res = fcl.db.WithContext(ctx).Order("created_at DESC").Limit(5).Find(&fallbackChatLog)
	} else {
		res = fcl.db.WithContext(ctx).Order("created_at DESC").Find(&fallbackChatLog)
	}

	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return fallbackChatLog, nil
}
