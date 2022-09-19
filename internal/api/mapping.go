package api

import (
	"github.com/alexeykirinyuk/go-product-api/internal/service/product/dto"
	pb "github.com/alexeykirinyuk/go-product-api/pkg/go-product-api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toProtobufProduct(product dto.Product) *pb.Product {
	return &pb.Product{
		Id:          product.ID,
		Name:        product.Name,
		Category:    product.Category,
		Description: product.Description,
		Brand:       product.Brand,
		Cost:        product.Cost,
		Currency:    pb.Currency(product.Currency),
		Created:     timestamppb.New(product.Created),
	}
}
