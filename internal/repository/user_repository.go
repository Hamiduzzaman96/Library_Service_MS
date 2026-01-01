package repository

import (
	"context"

	"github.com/Hamiduzzaman96/Library_Service_MS/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id int64) (*domain.User, error)
}
