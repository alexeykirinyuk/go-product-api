package model

import (
	vo "github.com/alexeykirinyuk/go-product-api/internal/service/value_objects"
	"time"
)

type Product struct {
	ID          uint64      `db:"id"`
	Name        string      `db:"name"`
	Category    string      `db:"category"`
	Description string      `db:"description"`
	Brand       string      `db:"brand"`
	Cost        float32     `db:"cost"`
	Currency    vo.Currency `db:"currency"`
	Created     time.Time   `db:"created"`
	Updated     time.Time   `db:"updated"`
}
