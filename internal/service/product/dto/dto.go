package dto

import (
	vo "github.com/alexeykirinyuk/go-product-api/internal/service/value_objects"
	"time"
)

type CreateProductCommand struct {
	Name        string
	Category    string
	Description string
	Brand       string
	Cost        float32
	Currency    vo.Currency
}

type Product struct {
	ID          uint64
	Name        string
	Category    string
	Description string
	Brand       string
	Cost        float32
	Currency    vo.Currency
	Created     time.Time
	Updated     time.Time
}
