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

func (a *actionHttpRepository) FindAll(ctx context.Context, name string) ([]*model.ActionHttp, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"name": name,
	})

	var actionHttps []*model.ActionHttp
	var res *gorm.DB
	if name == "" {
		res = a.db.WithContext(ctx).Find(&actionHttps)
	} else {
		res = a.db.WithContext(ctx).Where("name LIKE ?", "%"+name+"%").Find(&actionHttps)
	}
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return actionHttps, nil
}

func (a *actionHttpRepository) FindByID(ctx context.Context, id string) (*model.ActionHttp, error) {
	if id == "" {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	actionHttp := &model.ActionHttp{}
	err := a.db.WithContext(ctx).Where("id = ?", id).Take(actionHttp).Error
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
		"name",
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

func (a *actionHttpRepository) Delete(ctx context.Context, id string) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := a.db.WithContext(ctx).Where("id = ?", id).Delete(&model.ActionHttp{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (a *actionHttpRepository) CountAll(ctx context.Context) (int64, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})
	var count int64
	err := a.db.WithContext(ctx).Model(&model.ActionHttp{}).Count(&count).Error
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return count, nil
}
