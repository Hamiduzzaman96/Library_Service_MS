package client

import (
	"context"
	"time"

	"github.com/Hamiduzzaman96/Library_Service_MS/proto/bookpb"
	"google.golang.org/grpc"
)

type BookServiceClient struct {
	client bookpb.BookServiceClient
}

func NewBookServiceClient(add string) (*BookServiceClient, error) {
	conn, err := grpc.Dial(
		add,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(3*time.Second),
	)

	if err != nil {
		return nil, err
	}
	return &BookServiceClient{client: bookpb.NewBookServiceClient(conn)}, nil
}

func (b *BookServiceClient) CheckBookExists(ctx context.Context, bookID int64) (bool, error) {
	resp, err := b.client.CheckBookExists(ctx, &bookpb.CheckBookExistsRequest{
		Id: bookID,
	})
	if err != nil {
		return false, err
	}
	return resp.Exists, nil
}
