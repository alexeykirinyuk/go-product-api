package product

import (
	"context"
	"github.com/alexeykirinyuk/go-product-api/internal/service/product/dto"
)

func (s *Service) GetProductById(ctx context.Context, id uint64) (dto.Product, error) {
	prod, err := s.repo.Get(ctx, id)
	if err != nil {
		return dto.Product{}, err
	}

	return toProductDto(prod), nil
}
