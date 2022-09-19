package api

import (
	"github.com/alexeykirinyuk/go-product-api/internal/service/product"
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
	s *product.Service
}

func NewProductAPI(s *product.Service) *ProductAPI {
	return &ProductAPI{s: s}
}
