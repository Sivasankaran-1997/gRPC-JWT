package main

import (
	pb "gRPC_jwt/users/proto"
	"gRPC_jwt/users/server/controller"
	"gRPC_jwt/users/server/database"
	"gRPC_jwt/users/server/respository"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen. %v", err)
	}

	database.Connect()
	database.Instance.AutoMigrate(&respository.User{})
	opts := []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)

	pb.RegisterUserServiceServer(srv, controller.NewUserControllerServer())

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve. %v", err)

	}

}
