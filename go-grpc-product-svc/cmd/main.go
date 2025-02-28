package main

import (
	"fmt"
	"go-grpc-product-svc/pkg/config"
	"go-grpc-product-svc/pkg/db"
	"go-grpc-product-svc/pkg/pb"
	"go-grpc-product-svc/pkg/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing", err)
	}
	fmt.Println("Product Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("failed to serve", err)
	}
}
