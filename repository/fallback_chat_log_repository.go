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

func (fcl *fallbackChatLogRepository) FindAllGroupCluster(ctx context.Context) ([]*model.ClusterData, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	var result []*model.ClusterData
	fcl.db.WithContext(ctx).Table("fallback_chat_logs").
		Select("cluster, id, chat, created_at, deleted_at").
		Group("cluster").
		Scan(&result)

	for i := range result {
		if result[i].Cluster.Valid {
			var logs []*model.FallbackChatLog
			fcl.db.WithContext(ctx).Table("fallback_chat_logs").
				Where("cluster = ?", result[i].Cluster).
				Find(&logs)
			result[i].Data = logs
		} else {
			var logs []*model.FallbackChatLog
			fcl.db.WithContext(ctx).Table("fallback_chat_logs").
				Where("cluster IS NULL").
				Find(&logs)
			result[i].Data = logs
		}
	}

	if fcl.db.Error != nil {
		log.Error(fcl.db.Error)
		return nil, fcl.db.Error
	}

	return result, nil
}

func (fcl *fallbackChatLogRepository) FindNullAndCluster(ctx context.Context) (*model.ResponseFallback, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	var fallbackChatLogs *model.ResponseFallback
	var res *gorm.DB

	var fallbackChatOldLogs []*model.FallbackChatLog
	resOld := fcl.db.WithContext(ctx).Order("created_at DESC").Find(&fallbackChatOldLogs)
	if resOld.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	var fallbackChatNewLogs []*model.FallbackChatLog
	resNew := fcl.db.WithContext(ctx).Where("cluster IS NULL").Order("created_at DESC").Find(&fallbackChatNewLogs)
	if resNew.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	fallbackChatLogs = &model.ResponseFallback{
		ExistingLog: fallbackChatOldLogs,
		NewLog:      fallbackChatNewLogs,
	}

	return fallbackChatLogs, nil
}

func (fcl *fallbackChatLogRepository) UpdateGroupCluster(ctx context.Context, fallback *model.FallbackChatLog, tx *gorm.DB) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":             ctx,
		"fallbackChatLog": fallback,
	})

	err := tx.WithContext(ctx).Select(
		"cluster",
	).Updates(fallback).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
