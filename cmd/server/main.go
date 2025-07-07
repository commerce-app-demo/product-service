package main

import (
	"net"
	"os"

	"github.com/commerce-app-demo/product-service/internal/server"

	productspb "github.com/commerce-app-demo/product-service/proto"
	"google.golang.org/grpc"
)

func main() {
	p := ":50051"
	l, err := net.Listen("tcp", p)

	if err != nil {
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	productspb.RegisterProductServiceServer(grpcServer, &server.ProductServiceServer{})

	grpcServer.Serve(l)
}
