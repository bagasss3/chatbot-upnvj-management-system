package service

import (
	"bytes"
	"cbupnvj/config"
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"cbupnvj/repository"
	"context"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v4"
)

type fallbackChatLogService struct {
	fallbackChatLogRepository model.FallbackChatLogRepository
	gormTransactioner         repository.GormTransactioner
}

func NewFallbackChatLogService(fallbackChatLogRepository model.FallbackChatLogRepository, gormTransactioner repository.GormTransactioner) model.FallbackChatLogService {
	return &fallbackChatLogService{
		fallbackChatLogRepository: fallbackChatLogRepository,
		gormTransactioner:         gormTransactioner,
	}
}

func (fcl *fallbackChatLogService) CreateFallbackChatLog(ctx context.Context, req model.CreateFallbackChatLogRequest) (*model.FallbackChatLog, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, constant.HttpValidationOrInternalErr(err)
	}

	newFallbackChatLogIntent := &model.FallbackChatLog{
		Id:      helper.GenerateID(),
		Chat:    req.Chat,
		Cluster: null.NewInt(0, false),
	}

	err := fcl.fallbackChatLogRepository.Create(ctx, newFallbackChatLogIntent)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return newFallbackChatLogIntent, nil
}

func (fcl *fallbackChatLogService) FindAllFallbackChatLog(ctx context.Context, page string) ([]*model.FallbackChatLog, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	if page != string(model.FallbackPageDashboard) && page != string(model.FallbackPageLog) {
		log.Error(constant.ErrInvalidArgument)
		return nil, constant.ErrInvalidArgument
	}

	fallbacks, err := fcl.fallbackChatLogRepository.FindAll(ctx, page)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return fallbacks, nil
}

func (fcl *fallbackChatLogService) FindAllFallbackChatLogGroupCluster(ctx context.Context) ([]*model.ClusterData, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	fallbacks, err := fcl.fallbackChatLogRepository.FindAllGroupCluster(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return fallbacks, nil
}

func (fcl *fallbackChatLogService) FindAllFallbackChatLogOldAndNew(ctx context.Context) (*model.ResponseFallback, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	fallbacks, err := fcl.fallbackChatLogRepository.FindNullAndCluster(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return fallbacks, nil
}

func (fcl *fallbackChatLogService) UpdateGroupCluster(ctx context.Context) ([]*model.ClusterData, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	fallbacks, err := fcl.fallbackChatLogRepository.FindNullAndCluster(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	jsonBody, err := json.Marshal(fallbacks)
	if err != nil {
		// handle error
		log.Error(err)
		return nil, constant.ErrInternal
	}

	putReq, err := http.NewRequest("POST", config.GroupFallbackLog(), bytes.NewBuffer([]byte(jsonBody)))

	if err != nil {
		log.Error(err)
	}

	putReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	putResp, err := client.Do(putReq)
	if err != nil {
		log.Error(err)
	}
	defer putResp.Body.Close()

	if putResp.StatusCode != http.StatusOK {
		// Handle the non-OK response here if needed
		log.Error(putResp.Status)
		return nil, constant.ErrInternal
	}

	var clusterData []*model.FallbackChatLog
	err = json.NewDecoder(putResp.Body).Decode(&clusterData)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	tx := fcl.gormTransactioner.Begin(ctx)
	for i := range clusterData {
		err := fcl.fallbackChatLogRepository.UpdateGroupCluster(ctx, clusterData[i], tx)
		if err != nil {
			log.Error(err)
			fcl.gormTransactioner.Rollback(tx)
			return nil, err
		}
	}

	if err = fcl.gormTransactioner.Commit(tx); err != nil {
		log.Error(err)
		fcl.gormTransactioner.Rollback(tx)
		return nil, err
	}

	resp, err := fcl.fallbackChatLogRepository.FindAllGroupCluster(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return resp, nil
}
