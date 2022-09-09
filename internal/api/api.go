package api

import (
	"github.com/alexeykirinyuk/go-product-api/internal/service/product_service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	pb "github.com/alexeykirinyuk/go-product-api/pkg/go-product-api"
)

var (
	totalProductNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "go_product_api_not_found_total",
		Help: "Total number of Products that were not found",
	})
)

type ProductAPI struct {
	pb.UnimplementedGoProductApiServiceServer
	s *product_service.ProductService
}

func NewProductAPI(s *product_service.ProductService) *ProductAPI {
	return &ProductAPI{s: s}
}
