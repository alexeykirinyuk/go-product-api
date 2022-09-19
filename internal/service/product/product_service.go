package product

import (
	"context"
	"database/sql"
	"github.com/alexeykirinyuk/go-product-api/internal/service/database"
	"github.com/alexeykirinyuk/go-product-api/internal/service/product/dto"
	"github.com/alexeykirinyuk/go-product-api/internal/service/product/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"time"
)

type ProductRepo interface {
	AddProduct(ctx context.Context, product *model.Product, tx *sqlx.Tx) error
	GetProduct(ctx context.Context, productID uint64, tx *sqlx.Tx) (model.Product, error)
	ListProducts(ctx context.Context) ([]model.Product, error)
	RemoveProduct(ctx context.Context, productID uint64, tx *sqlx.Tx) error
}

type EventRepo interface {
	AddEvent(ctx context.Context, event *model.ProductEvent, tx *sqlx.Tx) error
}

type Service struct {
	productRepo ProductRepo
	eventRepo   EventRepo
	db          *sqlx.DB
}

func NewService(
	productRepo ProductRepo,
	eventRepo EventRepo,
	db *sqlx.DB) *Service {
	return &Service{
		productRepo: productRepo,
		eventRepo:   eventRepo,
		db:          db,
	}
}

func (s *Service) CreateProduct(
	ctx context.Context,
	productDto dto.CreateProductCommand,
) (dto.Product, error) {
	var product model.Product

	txErr := database.WithTx(ctx, s.db, sql.LevelReadCommitted, func(ctx context.Context, tx *sqlx.Tx) error {
		product = model.Product{
			Name:        productDto.Name,
			Category:    productDto.Category,
			Description: productDto.Description,
			Brand:       productDto.Brand,
			Cost:        productDto.Cost,
			Currency:    productDto.Currency,
			Created:     time.Now().UTC(),
		}

		if err := s.productRepo.AddProduct(ctx, &product, nil); err != nil {
			return errors.Wrap(err, "s.productRepo.AddProduct()")
		}

		event := &model.ProductEvent{
			ProductID: product.ID,
			Type:      model.EventTypeProductCreated,
			Status:    model.EventStatusCreated,
			Updated:   time.Now().UTC(),
			Payload: model.Payload{
				Created: &model.ProductCreatedPayload{
					ID:          product.ID,
					Name:        product.Name,
					Category:    product.Category,
					Description: product.Description,
					Brand:       product.Brand,
					Cost:        product.Cost,
					Currency:    product.Currency,
				},
			},
		}

		if err := s.eventRepo.AddEvent(ctx, event, tx); err != nil {
			return errors.Wrap(err, "eventRepo.AddEvent()")
		}

		return nil
	})
	if txErr != nil {
		return dto.Product{}, errors.Wrap(txErr, "transaction error")
	}

	return toProductDto(product), nil
}

func (s *Service) GetProductById(ctx context.Context, id uint64) (dto.Product, error) {
	prod, err := s.productRepo.GetProduct(ctx, id, nil)
	if err != nil {
		return dto.Product{}, err
	}

	return toProductDto(prod), nil
}

func (s *Service) GetProducts(ctx context.Context) ([]dto.Product, error) {
	products, err := s.productRepo.ListProducts(ctx)
	if err != nil {
		return []dto.Product{}, err
	}

	res := make([]dto.Product, 0, len(products))
	for _, p := range products {
		res = append(res, toProductDto(p))
	}

	return res, nil
}

func (s *Service) DeleteProduct(ctx context.Context, productID uint64) (dto.Product, error) {
	var product model.Product

	txErr := database.WithTx(ctx, s.db, sql.LevelReadCommitted, func(ctx context.Context, tx *sqlx.Tx) error {
		var err error
		product, err = s.productRepo.GetProduct(ctx, productID, tx)
		if err != nil {
			return errors.Wrap(err, "s.productRepo.GetProduct()")
		}

		if err := s.productRepo.RemoveProduct(ctx, productID, tx); err != nil {
			return errors.Wrap(err, "s.productRepo.RemoveProduct()")
		}

		event := model.ProductEvent{
			ProductID: productID,
			Type:      model.EventTypeProductDeleted,
			Status:    model.EventStatusCreated,
			Updated:   time.Now().UTC(),
			Payload: model.Payload{
				Removed: &model.ProductRemovedPayload{ID: productID},
			},
		}

		if err := s.eventRepo.AddEvent(ctx, &event, tx); err != nil {
			return errors.Wrap(err, "s.eventRepo.AddEvent()")
		}

		return nil
	})
	if txErr != nil {
		return dto.Product{}, errors.Wrap(txErr, "transaction error")
	}

	return toProductDto(product), nil
}
