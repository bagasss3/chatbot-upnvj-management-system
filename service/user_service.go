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

func (u *userService) FindAllAdmin(ctx context.Context) ([]*model.User, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	users, err := u.userRepository.FindAll(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return users, nil
}

func (u *userService) FindAdminByID(ctx context.Context, id int64) (*model.User, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	if id <= 0 {
		return nil, constant.ErrInvalidArgument
	}

	user, err := u.userRepository.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if user == nil {
		return nil, constant.ErrNotFound
	}

	return user, nil
}

func (u *userService) UpdateAdmin(ctx context.Context, id int64, req model.UpdateAdminRequest) (*model.User, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
		"req": req,
	})

	if req.Password != req.Repassword {
		log.Error("Password mismatch")
		return nil, constant.ErrPasswordMismatch
	}

	user, err := u.FindAdminByID(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	newPlainPassword := helper.GeneratePassword()
	newHashedPassword, err := helper.HashString(newPlainPassword)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	user.Name = req.Name
	user.MajorId = req.MajorId
	user.Password = newHashedPassword

	err = u.userRepository.Update(ctx, user.Id, user)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return user, nil
}

func (u *userService) DeleteAdminByID(ctx context.Context, id int64) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	_, err := u.FindAdminByID(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	err = u.userRepository.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}

func isValidUserType(t model.UserType) bool {
	return t == model.UserAdmin || t == model.UserSuperAdmin
}
