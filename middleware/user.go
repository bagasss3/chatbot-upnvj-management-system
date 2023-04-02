package middleware

import (
	"cbupnvj/model"
	"context"
)

// NewUserFromSession return new user from session
func NewUserAuth(user int64, role model.UserType) model.UserAuth {
	return model.UserAuth{
		UserID: user,
		Role:   role,
	}
}

type contextKey string

// use module path to make it unique
const userCtxKey contextKey = "cbupnvj/middleware/user"

// SetUserToCtx set user to context
func SetUserToCtx(ctx context.Context, user model.UserAuth) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}

// GetUserFromCtx get user from context
func GetUserFromCtx(ctx context.Context) *model.UserAuth {
	user, ok := ctx.Value(userCtxKey).(model.UserAuth)
	if !ok {
		return nil
	}
	return &user
}
