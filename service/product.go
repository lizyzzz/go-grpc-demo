package service

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/protobuf/types/known/anypb"
)

func CreateProductService() *productService {
	return &productService{
		UnimplementedProdServiceServer: UnimplementedProdServiceServer{},
	}
}

// 实现接口(把 UnimplementedProdServiceServer 作为匿名成员, 自动获得其所有方法, 即使没有实现接口方法也不会错, 并且能自动适应前向版本)
type productService struct {
	UnimplementedProdServiceServer
}

// 1. 实现普通模式的服务
func (p *productService) GetProductStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	// 实现具体的业务逻辑

	// 假设这里做了查询
	stock := p.GetStockById(request.GetProdId())
	content := &Content{
		Msg: "hello any",
	}
	any, _ := anypb.New(content)
	prod_resp := &ProductResponse{
		ProdStock: stock,
		User: &User{
			Username:  "lizy",
			Age:       25,
			Password:  "123456",
			Addresses: []string{"shenzhen"},
		},
		Data: any,
	}
	return prod_resp, nil
}

func (p *productService) GetStockById(id int32) int32 {
	return id * 2
}

// 2. 实现客户端流模式的服务
func (p *productService) UpdateProductStockClientStream(stream ProdService_UpdateProductStockClientStreamServer) error {
	count := 0
	for {
		// 接受客户端发送过来的信息
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				stock := p.GetStockById(int32(count))
				content := &Content{
					Msg: "hello any",
				}
				any, _ := anypb.New(content)
				prod_resp := &ProductResponse{
					ProdStock: stock,
					User: &User{
						Username:  "lizy",
						Age:       25,
						Password:  "123456",
						Addresses: []string{"shenzhen"},
					},
					Data: any,
				}
				err2 := stream.SendAndClose(prod_resp)
				if err2 != nil {
					return err2
				}
				break
			}
			return err
		}
		count++
		fmt.Println("recv client prodId:", recv.GetProdId(), "count:", count)
	}
	return nil
}

// 3. 实现服务端流模式的服务
func (p *productService) GetProductStockServerStream(request *ProductRequest, stream ProdService_GetProductStockServerStreamServer) error {
	content := &Content{
		Msg: "hello any",
	}
	any, _ := anypb.New(content)
	// 发送十次
	for i := 0; i < 10; i++ {
		prod_resp := &ProductResponse{
			ProdStock: p.GetStockById(int32(i) + request.GetProdId()),
			User: &User{
				Username:  "lizy",
				Age:       25,
				Password:  "123456",
				Addresses: []string{"shenzhen"},
			},
			Data: any,
		}
		err := stream.Send(prod_resp)
		if err != nil {
			return err
		}
	}

	return nil
}

// 4. 实现双向流的服务
func (p *productService) SayHelloStream(server_stream ProdService_SayHelloStreamServer) error {
	for {
		// 接收消息
		clientMsg, err := server_stream.Recv()
		if err != nil {
			return err
		}
		// 逻辑处理, 发送消息
		time.Sleep(time.Second * 1)
		fmt.Println("stream recv clientMsg", clientMsg.String())
		serverMsg := &ServerMsg{
			Info:   clientMsg.GetInfo(),
			Result: clientMsg.GetFirst() + clientMsg.GetSecond(),
		}
		err = server_stream.Send(serverMsg)
		if err != nil {
			return err
		}
	}
}
