package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go-grpc-demo/service"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {
	// 读取证书(给定公钥和私钥)

	// 1. 单向证书认证
	// creds, err2 := credentials.NewServerTLSFromFile("cert/server.pem", "cert/server.key")
	// if err2 != nil {
	// 	panic(err2)
	// }
	// rpcServer := grpc.NewServer(grpc.Creds(creds))

	// 2. 证书认证-双向认证
	// 从证书相关文件中读取和解析信息, 得到证书公钥-私钥对
	cert, err := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	if err != nil {
		panic(err)
	}

	// 创建一个新的空的 CertPool
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile("cert/ca.crt")
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
		ClientAuth: tls.RequireAndVerifyClientCert,
		// 设置根证书的集合, 校验方式使用 ClientAuth 中设定的模式
		ClientCAs: certPool,
	})

	// // 使用证书创建 rpcServer
	// rpcServer := grpc.NewServer(grpc.Creds(creds))

	// 3. 使用 token 认证
	// 实现一个拦截器
	var authInterceptor grpc.UnaryServerInterceptor
	authInterceptor = func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		// 拦截普通方法请求, 验证 token
		err = Auth(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}

	rpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(authInterceptor))
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

func Auth(ctx context.Context) error {
	// 实际上就是拿到 传输过来的用户名和密码
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var user string
	var password string

	if val, ok := md["user"]; ok {
		user = val[0]
	}

	if val, ok := md["password"]; ok {
		password = val[0]
	}

	if user != "admin" || password != "123456" {
		return status.Errorf(codes.Unauthenticated, "token unavailable")
	}
	return nil
}
