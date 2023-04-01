package service

import (
	"cbupnvj/config"
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type authService struct {
	userRepository    model.UserRepository
	sessionRepository model.SessionRepository
}

func NewAuthService(userRepository model.UserRepository, sessionRepository model.SessionRepository) model.AuthService {
	return &authService{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}

func (a *authService) LoginByEmailAndPassword(ctx context.Context, req model.LoginRequest) (*model.Session, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if user == nil {
		log.Error("user not found")
		return nil, constant.ErrNotFound
	}

	checkPwd := helper.IsHashStringMatch([]byte(req.PlainPassword), []byte(user.Password))
	if !checkPwd {
		log.Error("wrong email / password")
		return nil, constant.ErrUnauthorized
	}

	accessToken, err := generateToken(user.Id, config.AccessTokenDuration())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	refreshToken, err := generateToken(user.Id, config.RefreshTokenDuration())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	now := time.Now()
	session := &model.Session{
		Id:                    helper.GenerateID(),
		AccessToken:           accessToken,
		AccessTokenExpiredAt:  now.Add(config.AccessTokenDuration()),
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: now.Add(config.RefreshTokenDuration()),
		UserID:                user.Id,
	}

	err = a.sessionRepository.Create(ctx, session)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return session, nil
}

func (a *authService) RefreshToken(ctx context.Context, req model.RefreshTokenRequest) (*model.Session, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	session, err := a.sessionRepository.FindByRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if session == nil {
		log.Error("token not found")
		return nil, constant.ErrNotFound
	}

	if session.RefreshTokenExpiredAt.Before(time.Now()) {
		log.Error("refresh token expired")
		return nil, constant.ErrRefreshTokenExpired
	}

	newAccessToken, err := generateToken(session.UserID, config.AccessTokenDuration())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	newRefreshToken, err := generateToken(session.UserID, config.RefreshTokenDuration())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	session.AccessToken = newAccessToken
	session.RefreshToken = newRefreshToken

	now := time.Now()
	session.AccessTokenExpiredAt = now.Add(config.AccessTokenDuration())
	session.RefreshTokenExpiredAt = now.Add(config.RefreshTokenDuration())

	err = a.sessionRepository.RefreshToken(ctx, session)
	if err != nil {
		log.Error(err)
	}

	return session, nil
}
