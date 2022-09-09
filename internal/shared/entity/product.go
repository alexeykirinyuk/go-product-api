package entity

import (
	"github.com/alexeykirinyuk/go-product-api/internal/shared/enum"
	"time"
)

type Product struct {
	ID          uint64        `db:"id"`
	Name        string        `db:"name"`
	Category    string        `db:"category"`
	Description string        `db:"description"`
	Brand       string        `db:"brand"`
	Cost        float32       `db:"cost"`
	Currency    enum.Currency `db:"currency"`
	Created     time.Time     `db:"created"`
}
