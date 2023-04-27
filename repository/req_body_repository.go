package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type reqBodyRepository struct {
	db *gorm.DB
}

func NewReqBodyRepository(db *gorm.DB) model.ReqBodyRepository {
	return &reqBodyRepository{
		db: db,
	}
}

func (r *reqBodyRepository) Create(ctx context.Context, reqBody *model.ReqBody) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"reqBody": reqBody,
	})

	err := r.db.WithContext(ctx).Create(reqBody).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (r *reqBodyRepository) FindAll(ctx context.Context, actionHttpID string) ([]*model.ReqBody, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":          ctx,
		"actionHttpID": actionHttpID,
	})

	var reqBodies []*model.ReqBody
	res := r.db.WithContext(ctx).Where("action_http_id = ?", actionHttpID).Find(&reqBodies)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return reqBodies, nil
}

func (r *reqBodyRepository) FindByID(ctx context.Context, id string) (*model.ReqBody, error) {
	if id == "" {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	reqBody := &model.ReqBody{}
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(reqBody).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return reqBody, nil
}

func (r *reqBodyRepository) Update(ctx context.Context, actionHttpID string, reqBody *model.ReqBody) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":          ctx,
		"actionHttpID": actionHttpID,
		"reqBody":      reqBody,
	})

	err := r.db.WithContext(ctx).Select(
		"req_name",
	).Updates(reqBody).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (r *reqBodyRepository) Delete(ctx context.Context, id string) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.ReqBody{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
