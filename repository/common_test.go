package repository

import (
	"cbupnvj/model/mock"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type repoTestKit struct {
	dbmock          sqlmock.Sqlmock
	db              *gorm.DB
	ctrl            *gomock.Controller
	mockUserRepo    *mock.MockUserRepository
	mockSessionRepo *mock.MockSessionRepository
}

func initializeRepoTestKit(t *testing.T) (kit *repoTestKit, close func()) {
	dbconn, dbmock, err := sqlmock.New()
	if err != nil {
		logrus.Fatal(err)
	}

	dbmock.ExpectQuery("^SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.0"))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: dbconn}), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	userRepo := mock.NewMockUserRepository(ctrl)
	sessionRepo := mock.NewMockSessionRepository(ctrl)

	tk := &repoTestKit{
		ctrl:            ctrl,
		dbmock:          dbmock,
		db:              gormDB,
		mockUserRepo:    userRepo,
		mockSessionRepo: sessionRepo,
	}

	close = func() {
		if conn, _ := tk.db.DB(); conn != nil {
			_ = conn.Close()
		}
	}

	return tk, close
}
