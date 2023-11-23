package main

import (
	"fmt"
	"go-grpc-demo/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	prod_service := service.CreateProductService()
	service.RegisterProdServiceServer(rpcServer, prod_service)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		panic(err)
	}

	fmt.Println("grpc Server start")
	err = rpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
