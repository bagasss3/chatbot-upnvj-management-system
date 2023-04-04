package service

import (
	"cbupnvj/config"
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"
	"fmt"
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

	checkPwd := helper.IsHashStringMatch([]byte(req.Password), []byte(user.Password))
	if !checkPwd {
		log.Error("wrong email / password")
		return nil, constant.ErrUnauthorized
	}

	userAuth := &model.UserAuth{
		UserID: user.Id,
		Role:   user.Type,
	}

	accessToken, err := generateToken(userAuth, model.ACCESS_TOKEN_TYPE, config.AccessTokenDuration())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	refreshToken, err := generateToken(userAuth, model.REFRESH_TOKEN_TYPE, config.RefreshTokenDuration())
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

	user, err := a.userRepository.FindByID(ctx, session.UserID)
	if err != nil {
		log.Error(err)
		return nil, constant.ErrNotFound
	}

	userAuth := &model.UserAuth{
		UserID: user.Id,
		Role:   user.Type,
	}

	newAccessToken, err := generateToken(userAuth, model.ACCESS_TOKEN_TYPE, config.AccessTokenDuration())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	newRefreshToken, err := generateToken(userAuth, model.REFRESH_TOKEN_TYPE, config.RefreshTokenDuration())
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

func (a *authService) ForgotPassword(ctx context.Context, req model.ForgotPasswordRequest) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		log.Error(err)
		return false, err
	}

	if user == nil {
		log.Error(err)
		return false, constant.ErrNotFound
	}

	newPlainPassword := helper.GeneratePassword()
	newHashedPassword, err := helper.HashString(newPlainPassword)
	if err != nil {
		log.Error(err)
		return false, err
	}

	userFieldsToUpdate := &model.User{
		Id:       user.Id,
		Password: newHashedPassword,
	}

	if err = a.userRepository.ResetPassword(ctx, userFieldsToUpdate); err != nil {
		log.Error(err)
		return false, err
	}

	sendMail(ctx, user.Email, "Forgot Password", fmt.Sprintf("Your new Password: %s", newPlainPassword))
	return true, nil
}
