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

func (r *ruleRepository) FindAll(ctx context.Context, name string) ([]*model.Rule, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"name": name,
	})

	var rules []*model.Rule
	var res *gorm.DB
	if name == "" {
		res = r.db.WithContext(ctx).Find(&rules)
	} else {
		res = r.db.WithContext(ctx).Where("rule_title LIKE ?", "%"+name+"%").Find(&rules)
	}
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return rules, nil
}

func (r *ruleRepository) FindByID(ctx context.Context, id string) (*model.Rule, error) {
	if id == "" {
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

func (r *ruleRepository) Update(ctx context.Context, id string, rule *model.Rule) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"id":   id,
		"rule": rule,
	})

	err := r.db.WithContext(ctx).Select(
		"intent_id",
		"data_id",
		"rule_title",
		"response_id",
		"type",
	).Updates(rule).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (r *ruleRepository) Delete(ctx context.Context, id string) error {
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

func (r *ruleRepository) CountAll(ctx context.Context) (int64, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Rule{}).Count(&count).Error
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return count, nil
}
