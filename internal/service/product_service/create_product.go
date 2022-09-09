package product_service

import (
	"context"
	"github.com/alexeykirinyuk/go-product-api/internal/service/dto"
	"github.com/alexeykirinyuk/go-product-api/internal/shared/entity"
	"time"
)

func (s *ProductService) CreateProduct(
	ctx context.Context,
	productDto dto.CreateProductCommand,
) (dto.Product, error) {
	product := entity.Product{
		Name:        productDto.Name,
		Category:    productDto.Category,
		Description: productDto.Description,
		Brand:       productDto.Brand,
		Cost:        productDto.Cost,
		Currency:    productDto.Currency,
		Created:     time.Now().UTC(),
	}

	if err := s.repo.Save(ctx, &product); err != nil {
		return dto.Product{}, err
	}

	return toProductDto(product), nil
}
