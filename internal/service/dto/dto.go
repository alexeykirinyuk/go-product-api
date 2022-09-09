package dto

import (
	"github.com/alexeykirinyuk/go-product-api/internal/shared/enum"
	"time"
)

type CreateProductCommand struct {
	Name        string
	Category    string
	Description string
	Brand       string
	Cost        float32
	Currency    enum.Currency
}

type Product struct {
	ID          uint64
	Name        string
	Category    string
	Description string
	Brand       string
	Cost        float32
	Currency    enum.Currency
	Created     time.Time
}
