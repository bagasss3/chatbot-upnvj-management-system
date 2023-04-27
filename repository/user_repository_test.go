package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
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
		mock.ExpectQuery("^SELECT .+ FROM \"users\"").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(123))

		res, err := repo.FindByID(ctx, "")
		require.NoError(t, err)
		require.NotNil(t, res)
	})
}
