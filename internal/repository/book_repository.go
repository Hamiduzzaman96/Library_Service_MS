package repository

import "github.com/Hamiduzzaman96/Library_Service_MS/internal/domain"

type BookRepository interface {
	Create(book *domain.Book) (int64, error)
	GetByID(id int64) (*domain.Book, error)
}
