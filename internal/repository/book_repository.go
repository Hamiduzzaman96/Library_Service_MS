package repository

import (
	"context"

	"github.com/Hamiduzzaman96/Library_Service_MS/internal/domain"
)

type BookRepository interface {
	Create(ctx context.Context, book *domain.Book) error
	GetByID(ctx context.Context, id int64) (*domain.Book, error)
}
