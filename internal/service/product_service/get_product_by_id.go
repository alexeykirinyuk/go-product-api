package product_service

import (
	"context"
	"github.com/alexeykirinyuk/go-product-api/internal/service/dto"
	"github.com/alexeykirinyuk/go-product-api/internal/shared/internal_errors"
)

func (s *ProductService) GetProductById(ctx context.Context, id uint64) (dto.Product, error) {
	prod, exists, err := s.repo.GetById(ctx, id)
	if err != nil {
		return dto.Product{}, err
	}

	if !exists {
		return dto.Product{}, internal_errors.ProductNotFound
	}

	return toProductDto(prod), nil
}
