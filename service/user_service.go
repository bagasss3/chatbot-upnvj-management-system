package service

import (
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
)

type userService struct {
	userRepository model.UserRepository
}

func NewUserService(userRepository model.UserRepository) model.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) CreateAdmin(ctx context.Context, req model.CreateAdminRequest) (*model.User, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"request": req,
	})

	if req.Password != req.Repassword {
		log.Error("Password mismatch")
		return nil, constant.ErrPasswordMismatch
	}

	if !isValidUserType(req.Type) {
		log.Error("Invalid user type")
		return nil, constant.ErrInvalidArgument
	}

	cipherPwd, err := helper.HashString(req.Password)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	user := &model.User{
		Id:       helper.GenerateID(),
		Name:     req.Name,
		Password: cipherPwd,
		Email:    req.Email,
		Type:     req.Type,
		MajorId:  req.MajorId,
	}

	err = u.userRepository.Create(ctx, user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return user, nil
}

func isValidUserType(t model.UserType) bool {
	return t == model.UserAdmin || t == model.UserSuperAdmin
}
