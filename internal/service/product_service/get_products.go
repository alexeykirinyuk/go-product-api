package product_service

import (
	"context"
	"github.com/alexeykirinyuk/go-product-api/internal/service/dto"
)

func (s *ProductService) GetProducts(ctx context.Context) ([]dto.Product, error) {
	products, err := s.repo.GetList(ctx)
	if err != nil {
		return []dto.Product{}, err
	}

	res := make([]dto.Product, 0, len(products))
	for _, p := range products {
		res = append(res, toProductDto(p))
	}

	return res, nil
}
