package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type ruleService struct {
	ruleRepository       model.RuleRepository
	intentRepository     model.IntentRepository
	actionHttpRepository model.ActionHttpRepository
	utteranceRepository  model.UtteranceRepository
}

func NewRuleService(ruleRepository model.RuleRepository, intentRepository model.IntentRepository, actionHttpRepository model.ActionHttpRepository, utteranceRepository model.UtteranceRepository) model.RuleService {
	return &ruleService{
		ruleRepository:       ruleRepository,
		intentRepository:     intentRepository,
		actionHttpRepository: actionHttpRepository,
		utteranceRepository:  utteranceRepository,
	}
}

func (r *ruleService) CreateRule(ctx context.Context, req model.CreateUpdateRuleRequest) (*model.Rule, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	err := r.ValidateRule(ctx, req.IntentId, req.ResponseId, req.Type)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	rule := &model.Rule{
		Id:         helper.GenerateID(),
		IntentId:   req.IntentId,
		ResponseId: req.ResponseId,
		Type:       req.Type,
		RuleTitle:  req.RuleTitle,
	}

	err = r.ruleRepository.Create(ctx, rule)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return rule, nil
}

func (r *ruleService) FindAllRule(ctx context.Context) ([]*model.Rule, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	rules, err := r.ruleRepository.FindAll(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return rules, nil
}

func (r *ruleService) FindRuleByID(ctx context.Context, id string) (*model.Rule, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	if id == "" {
		return nil, constant.ErrInvalidArgument
	}

	rule, err := r.ruleRepository.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if rule == nil {
		return nil, constant.ErrNotFound
	}

	return rule, nil
}

func (r *ruleService) UpdateRule(ctx context.Context, id string, req model.CreateUpdateRuleRequest) (*model.Rule, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
		"req": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	err := r.ValidateRule(ctx, req.IntentId, req.ResponseId, req.Type)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	rule, err := r.FindRuleByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	rule.IntentId = req.IntentId
	rule.ResponseId = req.ResponseId
	rule.RuleTitle = req.RuleTitle
	rule.Type = req.Type

	err = r.ruleRepository.Update(ctx, id, rule)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return rule, nil
}

func (r *ruleService) DeleteRule(ctx context.Context, id string) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	_, err := r.FindRuleByID(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	err = r.ruleRepository.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}

func (r *ruleService) ValidateRule(ctx context.Context, intentId, dataId string, ruleType model.RuleType) error {
	if ruleType != model.RuleAction && ruleType != model.RuleUtterance {
		return constant.ErrInvalidArgument
	}

	intent, err := r.intentRepository.FindByID(ctx, intentId)
	if err != nil {
		return err
	}

	if intent == nil {
		return constant.ErrNotFound
	}

	if ruleType == model.RuleAction {
		action, err := r.actionHttpRepository.FindByID(ctx, dataId)
		if err != nil {
			return err
		}

		if action == nil {
			return constant.ErrNotFound
		}
	} else {
		utterance, err := r.utteranceRepository.FindByID(ctx, dataId)
		if err != nil {
			return err
		}

		if utterance == nil {
			return constant.ErrNotFound
		}
	}

	return nil
}
