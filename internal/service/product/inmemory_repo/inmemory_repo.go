package inmemory_repo

import (
	"context"
	"github.com/alexeykirinyuk/go-product-api/internal/service/product"
	"sync"
)

type ProductRepository struct {
	products map[uint64]product.Product
	nextID   uint64
	mtx      sync.RWMutex
}

// NewRepo returns ProductRepository struct
func NewRepo() *ProductRepository {
	return &ProductRepository{
		products: map[uint64]product.Product{},
		mtx:      sync.RWMutex{},
		nextID:   1,
	}
}

func (r *ProductRepository) Add(_ context.Context, product *product.Product) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if product.ID == 0 {
		product.ID = r.nextID
		r.nextID++
	}

	r.products[product.ID] = *product

	return nil
}

func (r *ProductRepository) Get(_ context.Context, productID uint64) (product.Product, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	p, ok := r.products[productID]
	if !ok {
		return product.Product{}, product.ProductNotFound
	}

	return p, nil
}

func (r *ProductRepository) List(_ context.Context) ([]product.Product, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	result := make([]product.Product, 0, len(r.products))
	for _, prod := range r.products {
		result = append(result, prod)
	}

	return result, nil
}

func (r *ProductRepository) Remove(_ context.Context, productID uint64) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	delete(r.products, productID)

	return nil
}
