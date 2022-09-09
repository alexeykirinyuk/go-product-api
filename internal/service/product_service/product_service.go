package product_service

import (
	"context"
	"github.com/alexeykirinyuk/go-product-api/internal/shared/entity"
)

type IProductRepository interface {
	Save(_ context.Context, product *entity.Product) error
	GetById(_ context.Context, productID uint64) (entity.Product, bool, error)
	GetList(_ context.Context) ([]entity.Product, error)
	Remove(_ context.Context, productID uint64) error
}

type ProductService struct {
	repo IProductRepository
}

func New(repo IProductRepository) *ProductService {
	return &ProductService{repo: repo}
}
