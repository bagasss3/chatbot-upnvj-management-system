package model

import (
	"github.com/kamva/mgm/v3"
)

type UserType string

const (
	UserAdmin      UserType = "ADMIN"
	UserSuperAdmin UserType = "SUPER_ADMIN"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string   `json:"email" bson:"email"`
	Password         string   `json:"password" bson:"password"`
	Name             string   `json:"name" bson:"name"`
	Type             UserType `json:"type" bson:"type"`
	Session          *Session `json:"session" bson:"-"`
}
