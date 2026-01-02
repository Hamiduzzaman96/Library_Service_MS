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
	"github.com/Hamiduzzaman96/Library_Service_MS/pkg/config"
	pb "github.com/Hamiduzzaman96/Library_Service_MS/proto/bookpb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	// db, _ := sql.Open("mysql", "user:pass@tcp(localhost:3306)/library")
	cfg := config.LoadBookConfig()

	db, err := sql.Open("mysql", cfg.DBDSN)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close()

	repo := mysql.NewBookMySQLRepository(db)
	u := usecase.NewBookUsecase(repo)

	// gRPC server
	handler := grpct.NewBookHandler(u)

	lis, err := net.Listen("tcp", cfg.GRPCPort)

	grpcServer := grpc.NewServer()
	pb.RegisterBookServiceServer(grpcServer, handler)

	go func() {
		log.Printf("Book gRPC server running at %s", cfg.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	// HTTP client-facing gateway (grpc → HTTP)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	err = pb.RegisterBookServiceHandlerFromEndpoint(ctx, mux,
		"localhost:50051",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP client-facing Book API running on :8080")
	http.ListenAndServe(":8080", mux)
}
