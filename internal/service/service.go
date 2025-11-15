package service

import (
	"context"
	"ozzus/auth-repository/internal/domain"
)

type AuthService interface {
	Register(ctx context.Context, email, password string) (*domain.User,error)
	Login(ctx context.Context, email, password string) ()
}