package repo

import (
	"context"
	"github.com/alexeykirinyuk/go-product-api/internal/service/database"
	"github.com/alexeykirinyuk/go-product-api/internal/service/product/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type EventRepo struct {
	db *sqlx.DB
}

func NewEventRepo(db *sqlx.DB) *EventRepo {
	return &EventRepo{db: db}
}

func (e *EventRepo) AddEvent(ctx context.Context, event *model.ProductEvent, tx *sqlx.Tx) error {
	query, args, err := database.StatementBuilder().
		Insert("event").
		Columns("product_id", "type", "status", "payload", "updated").
		Values(event.ProductID, event.Type, event.Status, event.Payload, event.Updated).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return errors.Wrap(err, "Create query error")
	}

	var getContexter database.GetContexter
	if tx == nil {
		getContexter = e.db
	} else {
		getContexter = tx
	}

	var id uint64
	if err := getContexter.GetContext(ctx, &id, query, args...); err != nil {
		return errors.Wrap(err, "GetContext error")
	}

	return nil
}
