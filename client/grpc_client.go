package main

import (
	"context"
	"fmt"
	"go-grpc-demo/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	// 一般使用 https
	// 输入公钥
	creds, err2 := credentials.NewClientTLSFromFile("../cert/server.pem", "*.lizyzzz.com")
	if err2 != nil {
		panic(err2)
	}

	// 无认证 grpc http/2
	// conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(creds))

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	prodClient := service.NewProdServiceClient(conn)

	request := &service.ProductRequest{
		ProdId: 123,
	}

	stockResponse, err := prodClient.GetProductStock(context.Background(), request)
	if err != nil {
		panic(err)
	}

	fmt.Println("Get stock", stockResponse.GetProdStock())

}
