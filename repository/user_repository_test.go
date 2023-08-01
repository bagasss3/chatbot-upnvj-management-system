package repository

import (
	"cbupnvj/model"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestUserRepository_FindByID(t *testing.T) {
	kit, closer := initializeRepoTestKit(t)
	defer closer()
	mock := kit.dbmock

	repo := userRepository{
		db: kit.db,
	}

	ctx := context.TODO()

	t.Run("ok - retrieve from db", func(t *testing.T) {
		mock.ExpectQuery("^SELECT \\* FROM `users` WHERE id = \\? AND `users`.`deleted_at` IS NULL LIMIT 1").
			WithArgs("123").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("123"))

		res, err := repo.FindByID(ctx, "123")
		require.NoError(t, err)
		require.NotNil(t, res)
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery("^SELECT \\* FROM `users` WHERE id = \\? AND `users`.`deleted_at` IS NULL LIMIT 1").
			WithArgs("123").
			WillReturnError(gorm.ErrRecordNotFound)

		res, err := repo.FindByID(ctx, "123")
		require.NoError(t, err)
		require.Nil(t, res)
	})
}

func TestUserRepository_Create(t *testing.T) {
	kit, closer := initializeRepoTestKit(t)
	defer closer()
	mock := kit.dbmock

	repo := userRepository{
		db: kit.db,
	}

	ctx := context.TODO()

	user := &model.User{
		Id:        "123",
		Email:     "test@gmail.com",
		Password:  "hash",
		Name:      "test",
		Type:      model.UserAdmin,
		MajorId:   "1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("ok", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("^INSERT INTO `users`").
			WithArgs(user.Id, user.Email, user.Password, user.Name, user.Type, user.MajorId, user.CreatedAt, user.UpdatedAt, user.DeletedAt).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		err := repo.Create(ctx, user)
		require.NoError(t, err)
	})

	t.Run("failed to insert user", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("^INSERT INTO `users`").
			WithArgs(user.Id, user.Email, user.Password, user.Name, user.Type, user.MajorId, user.CreatedAt, user.UpdatedAt, user.DeletedAt).
			WillReturnError(errors.New("db error"))
		err := repo.Create(ctx, user)
		require.Error(t, err)
	})
}

func TestUserRepository_Update(t *testing.T) {
	kit, closer := initializeRepoTestKit(t)
	defer closer()
	mock := kit.dbmock

	repo := userRepository{
		db: kit.db,
	}

	ctx := context.TODO()

	user := &model.User{
		Id:        "123",
		Email:     "test@gmail.com",
		Password:  "hash",
		Name:      "test",
		Type:      model.UserAdmin,
		MajorId:   "1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("ok", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("^UPDATE `users`").
			WithArgs(user.Password, user.Name, user.MajorId, sqlmock.AnyArg(), user.Id).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		err := repo.Update(ctx, user.Id, user)
		require.NoError(t, err)
	})

	t.Run("failed to update user", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("^UPDATE `users`").
			WithArgs(user.Password, user.Name, user.MajorId, sqlmock.AnyArg(), user.Id).
			WillReturnError(errors.New("db error"))
		err := repo.Update(ctx, user.Id, user)
		require.Error(t, err)
	})
}

func TestUserRepository_FindAll(t *testing.T) {
	kit, closer := initializeRepoTestKit(t)
	defer closer()
	mock := kit.dbmock

	repo := userRepository{
		db: kit.db,
	}

	ctx := context.TODO()

	t.Run("ok - found", func(t *testing.T) {
		ids := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		resRows := sqlmock.NewRows([]string{"id"})
		for _, v := range ids {
			resRows.AddRow(v)
		}
		mock.ExpectQuery("^SELECT \\* FROM `users` WHERE type = \\? AND `users`.`deleted_at` IS NULL ORDER BY created_at DESC").
			WithArgs(model.UserAdmin).
			WillReturnRows(resRows)

		res, err := repo.FindAll(ctx)
		require.NoError(t, err)
		require.NotNil(t, res)
	})

	t.Run("ok - not found", func(t *testing.T) {
		resRows := sqlmock.NewRows([]string{"id"})

		mock.ExpectQuery("^SELECT \\* FROM `users` WHERE type = \\? AND `users`.`deleted_at` IS NULL ORDER BY created_at DESC").
			WithArgs(model.UserAdmin).
			WillReturnRows(resRows)

		res, err := repo.FindAll(ctx)
		require.NoError(t, err)
		require.NotNil(t, res)
		assert.Equal(t, 0, len(res))
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery("^SELECT \\* FROM `users` WHERE type = \\? AND `users`.`deleted_at` IS NULL ORDER BY created_at DESC").
			WithArgs(model.UserAdmin).
			WillReturnError(gorm.ErrInvalidDB)

		res, err := repo.FindAll(ctx)
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestUserRepository_FindByEmail(t *testing.T) {
	kit, closer := initializeRepoTestKit(t)
	defer closer()
	mock := kit.dbmock

	repo := userRepository{
		db: kit.db,
	}

	ctx := context.TODO()

	user := &model.User{
		Id:    "123",
		Email: "test@gmail.com",
	}

	t.Run("ok - retrieve from db", func(t *testing.T) {
		mock.ExpectQuery("^SELECT `id` FROM `users` WHERE email = \\? AND `users`.`deleted_at` IS NULL LIMIT 1").
			WithArgs(user.Email).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(user.Id))
		mock.ExpectQuery("^SELECT \\* FROM `users` WHERE id = \\? AND `users`.`deleted_at` IS NULL LIMIT 1").
			WithArgs(user.Id).
			WillReturnRows(sqlmock.NewRows([]string{"id", "email"}).AddRow(user.Id, user.Email))
		res, err := repo.FindByEmail(ctx, user.Email)
		require.NoError(t, err)
		require.NotNil(t, res)
		assert.Equal(t, user, res)
	})

	t.Run("ok - not found", func(t *testing.T) {
		mock.ExpectQuery("^SELECT `id` FROM `users` WHERE email = \\? AND `users`.`deleted_at` IS NULL LIMIT 1").
			WithArgs(user.Email).
			WillReturnError(errors.New("error db"))

		res, err := repo.FindByEmail(ctx, user.Email)
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestUserRepository_ResetPassword(t *testing.T) {
	kit, closer := initializeRepoTestKit(t)
	defer closer()
	mock := kit.dbmock

	repo := userRepository{
		db: kit.db,
	}

	ctx := context.TODO()

	user := &model.User{
		Id:        "123",
		Email:     "test@gmail.com",
		Password:  "newhash",
		Name:      "test",
		Type:      model.UserAdmin,
		MajorId:   "1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("ok", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("^UPDATE `users`").
			WithArgs(user.Password, sqlmock.AnyArg(), user.Id).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		err := repo.ResetPassword(ctx, user)
		require.NoError(t, err)
	})

	t.Run("failed to update user", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("^UPDATE `users`").
			WithArgs(user.Password, sqlmock.AnyArg(), user.Id).
			WillReturnError(errors.New("db error"))
		err := repo.ResetPassword(ctx, user)
		require.Error(t, err)
	})
}

func TestUserRepository_Delete(t *testing.T) {
	kit, closer := initializeRepoTestKit(t)
	defer closer()
	mock := kit.dbmock

	repo := userRepository{
		db: kit.db,
	}

	ctx := context.TODO()

	user := &model.User{
		Id:        "123",
		Email:     "test@gmail.com",
		Password:  "newhash",
		Name:      "test",
		Type:      model.UserAdmin,
		MajorId:   "1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("ok", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("^UPDATE `users`").
			WithArgs(sqlmock.AnyArg(), user.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		err := repo.Delete(ctx, user.Id)
		require.NoError(t, err)
	})

	t.Run("failed to delete user", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("^UPDATE `users`").
			WithArgs(sqlmock.AnyArg(), user.Id).
			WillReturnError(errors.New("db error"))
		err := repo.Delete(ctx, user.Id)
		require.Error(t, err)
	})
}
