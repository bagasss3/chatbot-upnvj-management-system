package model

import (
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	mgm.DefaultModel      `bson:",inline"`
	AccessToken           string             `json:"access_token" bson:"access_token"`
	RefreshToken          string             `json:"refresh_token" bson:"refresh_token"`
	RefreshTokenExpiredAt time.Time          `json:"refresh_token_expired_at" bson:"refresh_token_expired_at"`
	UserID                primitive.ObjectID `json:"user_id" bson:"user_id"`
}
