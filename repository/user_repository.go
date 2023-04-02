package repository

import (
	"cbupnvj/model"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(ctx context.Context, user *model.User) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"user": user,
	})

	err := u.db.WithContext(ctx).Create(user).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (u *userRepository) FindByID(ctx context.Context, id int64) (*model.User, error) {
	if id <= 0 {
		return nil, nil
	}

	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	user := &model.User{}
	err := u.db.WithContext(ctx).Take(user, "id = ?", id).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

	return user, nil
}

func (u *userRepository) FindByEmail(ctx context.Context, userEmail string) (*model.User, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"userEmail": userEmail,
	})

	id := int64(0)
	err := u.db.Model(model.User{}).Select("id").Take(&id, "email = ?", userEmail).Error
	switch err {
	case nil:
		return u.FindByID(ctx, id)
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}
}

func (u *userRepository) ResetPassword(ctx context.Context, user *model.User) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"user": user,
	})

	err := u.db.WithContext(ctx).Select(
		"password",
	).Updates(user).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (u *userRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	var users []*model.User
	res := u.db.WithContext(ctx).Where("type = ?", "ADMIN").Find(&users)
	if res.Error != nil {
		log.Error(res.Error)
		return users, res.Error
	}

	return users, nil
}

func (u *userRepository) Update(ctx context.Context, id int64, user *model.User) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":  ctx,
		"id":   id,
		"user": user,
	})

	err := u.db.WithContext(ctx).Select(
		"name",
		"major_id",
		"password",
	).Updates(user).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (u *userRepository) Delete(ctx context.Context, id int64) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := u.db.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
