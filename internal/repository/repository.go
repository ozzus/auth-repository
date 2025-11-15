package repository

import (
	"context"
	"ozzus/auth-repository/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	IsAdmin(ctx context.Context, userID string) (bool, error)
}

type TokenRepository interface {
	SaveRefreshToken(ctx context.Context, userID, refreshToken string) error
	GetUserIDByRefreshToken(ctx context.Context, refreshToken string) (string, error)
	DeleteRefreshToken(ctx context.Context, userID, refreshToken string) error
}
