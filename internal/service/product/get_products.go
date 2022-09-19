package product

import (
	"context"
	"github.com/alexeykirinyuk/go-product-api/internal/service/product/dto"
)

func (s *Service) GetProducts(ctx context.Context) ([]dto.Product, error) {
	products, err := s.repo.List(ctx)
	if err != nil {
		return []dto.Product{}, err
	}

	res := make([]dto.Product, 0, len(products))
	for _, p := range products {
		res = append(res, toProductDto(p))
	}

	return res, nil
}
