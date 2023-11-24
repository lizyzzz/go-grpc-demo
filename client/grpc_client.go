package main

import (
	"context"
	"fmt"
	"go-grpc-demo/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 无认证 grpc http/2
	// 但一般使用 https
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))

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
