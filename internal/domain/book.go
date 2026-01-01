package domain

import "context"

type Book struct {
	ID     int64
	Title  string
	Author string
}

type BookServiceClient interface {
	CheckBookExists(ctx context.Context, bookID int64) (bool, error)
}
