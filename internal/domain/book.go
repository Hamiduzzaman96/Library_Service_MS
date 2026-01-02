package domain

import (
	"context"

	"github.com/Hamiduzzaman96/Library_Service_MS/proto/bookpb"
)

type Book struct {
	ID     int64
	Title  string
	Author string
}

type BookServiceClient interface {
	CheckBookExists(ctx context.Context, bookID int64) (bool, error)
}

type BookGRPCClient struct {
	client bookpb.BookServiceClient
}

// NewBookGRPCClient creates a new adapter
func NewBookGRPCClient(client bookpb.BookServiceClient) *BookGRPCClient {
	return &BookGRPCClient{client: client}
}

// CheckBookExists calls gRPC Book service and adapts the response
func (b *BookGRPCClient) CheckBookExists(ctx context.Context, bookID int64) (bool, error) {
	resp, err := b.client.GetBook(ctx, &bookpb.GetBookRequest{
		Id: bookID,
	})
	if err != nil {
		return false, err
	}

	if resp == nil || resp.Id == 0 {
		return false, nil
	}

	return true, nil
}
