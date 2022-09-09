package product_service

import (
	"context"
	"github.com/alexeykirinyuk/go-product-api/internal/service/dto"
	"github.com/alexeykirinyuk/go-product-api/internal/shared/internal_errors"
)

func (s *ProductService) DeleteProduct(ctx context.Context, productID uint64) (dto.Product, error) {
	prod, exists, err := s.repo.GetById(ctx, productID)
	if err != nil {
		return dto.Product{}, err
	}

	if !exists {
		return dto.Product{}, internal_errors.ProductNotFound
	}

	err = s.repo.Remove(ctx, productID)
	if err != nil {
		return dto.Product{}, err
	}

	return toProductDto(prod), nil
}
