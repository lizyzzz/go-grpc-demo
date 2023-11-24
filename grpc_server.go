package main

import (
	"fmt"
	"go-grpc-demo/service"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// 读取证书(给定公钥和私钥)
	cerds, err2 := credentials.NewServerTLSFromFile("cert/server.pem", "cert/server.key")
	if err2 != nil {
		panic(err2)
	}
	// 使用证书创建 rpcServer
	rpcServer := grpc.NewServer(grpc.Creds(cerds))
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
