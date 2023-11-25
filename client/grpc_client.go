package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go-grpc-demo/client/auth"
	"go-grpc-demo/service"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	// 一般使用 https

	// 1. 无认证 grpc http/2
	// conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))

	// 2. 单向认证
	// 输入公钥
	// creds, err2 := credentials.NewClientTLSFromFile("../cert/server.pem", "*.lizyzzz.com")
	// if err2 != nil {
	// 	panic(err2)
	// }

	// 3. 证书认证-双向认证
	// 从证书相关文件中读取和解析信息, 得到证书公钥-私钥对
	cert, err := tls.LoadX509KeyPair("../cert/client.pem", "../cert/client.key")
	if err != nil {
		panic(err)
	}

	// 创建一个新的空的 CertPool
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile("../cert/ca.crt")
	if err != nil {
		panic(err)
	}

	// 尝试解析所传入的 PEM 编码的证书, 如果解析成功会将其加到 CertPool 中, 便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链,允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书, 可以根据实际情况选用以下参数
		ServerName: "*.lizyzzz.com",
		// 设置根证书的集合, 校验方式使用 ClientAuth 中设定的模式
		RootCAs: certPool,
	})

	// 4. 使用 token 认证
	token := &auth.Authentication{
		User:     "admin",
		Password: "123456",
	}
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(token))

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
	fmt.Println("Get User", stockResponse.GetUser().String())
	fmt.Println("Get User", stockResponse.GetData().String())
}
