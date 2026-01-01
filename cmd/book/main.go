package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/Hamiduzzaman96/Library_Service_MS/internal/handler/grpct"
	"github.com/Hamiduzzaman96/Library_Service_MS/internal/repository/mysql"
	"github.com/Hamiduzzaman96/Library_Service_MS/internal/usecase"
	pb "github.com/Hamiduzzaman96/Library_Service_MS/proto/bookpb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	db, _ := sql.Open("mysql", "user:pass@tcp(localhost:3306)/library")

	repo := mysql.NewBookMySQLRepository(db)
	u := usecase.NewBookUsecase(repo)

	// gRPC server
	handler := grpct.NewBookHandler(u)
	lis, _ := net.Listen("tcp", ":50051")
	grpcServer := grpc.NewServer()
	pb.RegisterBookServiceServer(grpcServer, handler)
	go grpcServer.Serve(lis)
	log.Println(" gRPC Book service running on :50051")

	// HTTP client-facing gateway (grpc → HTTP)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	err := pb.RegisterBookServiceHandlerFromEndpoint(ctx, mux,
		"localhost:50051",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP client-facing Book API running on :8080")
	http.ListenAndServe(":8080", mux)
}
