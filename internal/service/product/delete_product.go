package product

import (
	"context"
	"github.com/alexeykirinyuk/go-product-api/internal/service/product/dto"
)

func (s *Service) DeleteProduct(ctx context.Context, productID uint64) (dto.Product, error) {
	prod, err := s.repo.Get(ctx, productID)
	if err != nil {
		return dto.Product{}, err
	}

	err = s.repo.Remove(ctx, productID)
	if err != nil {
		return dto.Product{}, err
	}

	return toProductDto(prod), nil
}
