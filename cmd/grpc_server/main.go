package main

import (
	"database/sql"
	"net"

	"github.com/davisande/grpc-study-project/internal/database"
	"github.com/davisande/grpc-study-project/internal/pb"
	"github.com/davisande/grpc-study-project/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoruDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoruDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
