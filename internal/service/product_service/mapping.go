package product_service

import (
	"github.com/alexeykirinyuk/go-product-api/internal/service/dto"
	"github.com/alexeykirinyuk/go-product-api/internal/shared/entity"
)

func toProductDto(product entity.Product) dto.Product {
	return dto.Product{
		ID:          product.ID,
		Name:        product.Name,
		Category:    product.Category,
		Description: product.Description,
		Brand:       product.Brand,
		Cost:        product.Cost,
		Currency:    product.Currency,
		Created:     product.Created,
	}
}
