package service

import (
	"context"

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

// func (p *productService) mustEmbedUnimplementedProdServiceServer() {}
