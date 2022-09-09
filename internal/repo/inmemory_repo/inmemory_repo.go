package inmemory_repo

import (
	"context"
	"github.com/alexeykirinyuk/go-product-api/internal/shared/entity"
	"sync"
)

type ProductRepository struct {
	products map[uint64]entity.Product
	nextID   uint64
	mtx      sync.RWMutex
}

// NewRepo returns ProductRepository struct
func NewRepo() *ProductRepository {
	return &ProductRepository{
		products: map[uint64]entity.Product{},
		mtx:      sync.RWMutex{},
		nextID:   1,
	}
}

func (r *ProductRepository) Save(_ context.Context, product *entity.Product) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if product.ID == 0 {
		product.ID = r.nextID
		r.nextID++
	}

	r.products[product.ID] = *product

	return nil
}

func (r *ProductRepository) GetById(_ context.Context, productID uint64) (entity.Product, bool, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	product, ok := r.products[productID]
	return product, ok, nil
}

func (r *ProductRepository) GetList(_ context.Context) ([]entity.Product, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	result := make([]entity.Product, 0, len(r.products))
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
