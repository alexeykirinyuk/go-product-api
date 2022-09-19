package product

import (
	"context"
	"github.com/alexeykirinyuk/go-product-api/internal/service/product/dto"
	"time"
)

func (s *Service) CreateProduct(
	ctx context.Context,
	productDto dto.CreateProductCommand,
) (dto.Product, error) {
	product := Product{
		Name:        productDto.Name,
		Category:    productDto.Category,
		Description: productDto.Description,
		Brand:       productDto.Brand,
		Cost:        productDto.Cost,
		Currency:    productDto.Currency,
		Created:     time.Now().UTC(),
	}

	if err := s.repo.Add(ctx, &product); err != nil {
		return dto.Product{}, err
	}

	return toProductDto(product), nil
}
