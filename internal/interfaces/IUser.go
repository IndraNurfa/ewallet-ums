package interfaces

import (
	"context"
	"ewallet-ums/internal/models"
	"time"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
	GetUserbyUsername(ctx context.Context, username string) (models.User, error)
	InsertNewUserSession(ctx context.Context, session *models.UserSession) error
	DeleteUserSession(ctx context.Context, token string) error
	GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error)
	UpdateTokenWByRefreshToken(ctx context.Context, token, refresh_token string, tokenExpired, updatedAt time.Time) error
	GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (models.UserSession, error)
}
