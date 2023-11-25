package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go-grpc-demo/client/auth"
	"go-grpc-demo/service"
	"io"
	"os"
	"time"

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
	// 1. 普通模式请求与响应
	stockResponse, err := prodClient.GetProductStock(context.Background(), request)
	if err != nil {
		panic(err)
	}

	fmt.Println("Get stock", stockResponse.GetProdStock())
	fmt.Println("Get User", stockResponse.GetUser().String())
	fmt.Println("Get Data", stockResponse.GetData().String())

	// 2. 客户端流模式的请求和获取响应
	fmt.Println("client stream ---------------- ")

	client_stream, err := prodClient.UpdateProductStockClientStream(context.Background())
	if err != nil {
		panic(err)
	}
	rsp := make(chan struct{}, 1)

	// 开启请求协程
	go func(stream service.ProdService_UpdateProductStockClientStreamClient, c chan struct{}) {
		for i := 0; i < 10; i++ {
			req := &service.ProductRequest{
				ProdId: 100,
			}
			err := stream.Send(req)
			if err != nil {
				panic(err)
			}
		}
		c <- struct{}{}
	}(client_stream, rsp)

	// 接收响应
	select {
	case <-rsp:
		recv, err := client_stream.CloseAndRecv()
		if err != nil {
			panic(err)
		}
		fmt.Println("Get stock", recv.GetProdStock())
		fmt.Println("Get User", recv.GetUser().String())
		fmt.Println("Get Data", recv.GetData().String())
	}

	// 3. 服务端流模式的请求和获取响应
	fmt.Println("server stream ---------------- ")
	server_stream, err := prodClient.GetProductStockServerStream(context.Background(), request)
	var count int32 = 0
	for {
		client_recv, err := server_stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("server stream recv end....")
				err = server_stream.CloseSend()
				if err != nil {
					panic(err)
				}
				break
			}
			panic(err)
		}
		count++
		fmt.Println("count:", count, "Get stock", client_recv.GetProdStock())
	}

	// 4. 双向流模式的请求和获取响应
	fmt.Println("client-server stream ---------------- ")
	stream, err := prodClient.SayHelloStream(context.Background())
	if err != nil {
		panic(err)
	}
	// 客户端发送和接收
	for i := 0; i < 10; i++ {
		// 先发送消息
		clientMsg := &service.ClientMsg{
			Info:   "Add operation",
			First:  int32(i),
			Second: 10,
		}
		stream.Send(clientMsg)
		fmt.Println("send clientMsg:", clientMsg.String())
		time.Sleep(time.Second * 1)
		// 再接收消息
		serverMsg, err := stream.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println("recv serverMsg:", serverMsg.String())
	}

}
