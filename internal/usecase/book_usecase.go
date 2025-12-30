package usecase

import (
	"context"

	"github.com/Hamiduzzaman96/Library_Service_MS/internal/domain"
	"github.com/Hamiduzzaman96/Library_Service_MS/internal/repository"
)

type BookUsecase struct {
	repo repository.BookRepository
}

func NewBookUsecase(repo repository.BookRepository) *BookUsecase {
	return &BookUsecase{repo: repo}
}
func (u *BookUsecase) CreateBook(ctx context.Context, book *domain.Book) error {
	return u.repo.Create(ctx, book)
}

func (u *BookUsecase) GetBook(ctx context.Context, id int64) (*domain.Book, error) {
	return u.repo.GetByID(ctx, id)
}
