package usecase

import (
	"context"

	"github.com/Hamiduzzaman96/Library_Service_MS/internal/domain"
	"github.com/Hamiduzzaman96/Library_Service_MS/internal/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: repo}
}

func (u *UserUsecase) Create(ctx context.Context, user *domain.User) error {
	return u.userRepo.Create(ctx, user)
}

func (u *UserUsecase) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return u.userRepo.GetByID(ctx, id)
}
