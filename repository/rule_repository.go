package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ruleRepository struct {
	db *gorm.DB
}

func NewRuleRepository(db *gorm.DB) model.RuleRepository {
	return &ruleRepository{
		db: db,
	}
}

func (r *ruleRepository) Create(ctx context.Context, rule *model.Rule) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"rule": rule,
	})

	err := r.db.WithContext(ctx).Create(rule).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (r *ruleRepository) FindAll(ctx context.Context) ([]*model.Rule, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	var rules []*model.Rule
	res := r.db.WithContext(ctx).Find(&rules)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return rules, nil
}

func (r *ruleRepository) FindByID(ctx context.Context, id int64) (*model.Rule, error) {
	if id <= 0 {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	rule := &model.Rule{}
	err := r.db.WithContext(ctx).Take(rule, "id = ?", id).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return rule, nil
}

func (r *ruleRepository) Update(ctx context.Context, id int64, rule *model.Rule) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"id":   id,
		"rule": rule,
	})

	err := r.db.WithContext(ctx).Select(
		"intent_id",
		"data_id",
		"rule_title",
		"type",
	).Updates(rule).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (r *ruleRepository) Delete(ctx context.Context, id int64) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Rule{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
