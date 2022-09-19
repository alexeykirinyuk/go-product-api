package product

import (
	"context"
)

type Repo interface {
	Add(ctx context.Context, product *Product) error
	Get(ctx context.Context, productID uint64) (Product, error)
	List(ctx context.Context) ([]Product, error)
	Remove(ctx context.Context, productID uint64) error
}

type Service struct {
	repo Repo
}

func New(repo Repo) *Service {
	return &Service{repo: repo}
}
