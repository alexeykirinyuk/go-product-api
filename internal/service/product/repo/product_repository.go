package product_repository

import (
	"context"
	"github.com/alexeykirinyuk/go-product-api/internal/service/product"
	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (p *ProductRepo) Add(ctx context.Context, product *product.Product) error {
	//database.StatementBuilder().
	//	Insert("products")
	panic("implement me")
}

func (p *ProductRepo) Get(ctx context.Context, productID uint64) (product.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductRepo) List(ctx context.Context) ([]product.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductRepo) Remove(ctx context.Context, productID uint64) error {
	//TODO implement me
	panic("implement me")
}
