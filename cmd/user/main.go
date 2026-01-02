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
	"github.com/Hamiduzzaman96/Library_Service_MS/proto/userpb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	//connect to MySQL
	db, err := sql.Open("mysql", "root:ramim#rimi96 @tcp(localhost:3306)/library")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect databse: %v", err)
	}
	defer db.Close()

	repo := mysql.NewUserMySQLRepository(db)
	u := usecase.NewUserUsecase(repo)

	// grpc server
	handler := grpct.NewUserHandler(u)
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, handler)

	go func() {
		log.Println("gRPC user service running on :50052")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	//HTTP client-facing (gRPC to HTTP)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()

	err = userpb.RegisterUserServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:50052",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP User API running on :8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal(err)
	}

}
