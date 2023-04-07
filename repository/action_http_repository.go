package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type actionHttpRepository struct {
	db *gorm.DB
}

func NewActionHttpRepository(db *gorm.DB) model.ActionHttpRepository {
	return &actionHttpRepository{
		db: db,
	}
}

func (a *actionHttpRepository) Create(ctx context.Context, actionHttp *model.ActionHttp) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":        ctx,
		"actionHttp": actionHttp,
	})

	err := a.db.WithContext(ctx).Create(actionHttp).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (a *actionHttpRepository) FindByID(ctx context.Context, actionId int64) (*model.ActionHttp, error) {
	if actionId <= 0 {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx":      ctx,
		"actionId": actionId,
	})

	actionHttp := &model.ActionHttp{}
	err := a.db.WithContext(ctx).Where("action_id = ?", actionId).Take(actionHttp).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return actionHttp, nil
}

func (a *actionHttpRepository) Update(ctx context.Context, actionHttp *model.ActionHttp) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":        ctx,
		"actionHttp": actionHttp,
	})

	err := a.db.WithContext(ctx).Select(
		"get_http_req",
		"post_http_req",
		"put_http_req",
		"del_http_req",
		"api_key",
		"text_response",
	).Updates(actionHttp).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (a *actionHttpRepository) Delete(ctx context.Context, actionId int64) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":      ctx,
		"actionId": actionId,
	})

	err := a.db.WithContext(ctx).Where("action_id = ?", actionId).Delete(&model.ActionHttp{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
