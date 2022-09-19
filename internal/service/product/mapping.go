package product

import (
	"github.com/alexeykirinyuk/go-product-api/internal/service/product/dto"
)

func toProductDto(product Product) dto.Product {
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
