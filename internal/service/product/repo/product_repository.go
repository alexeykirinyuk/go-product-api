package product_repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/alexeykirinyuk/go-product-api/internal/service/database"
	"github.com/alexeykirinyuk/go-product-api/internal/service/product"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type ProductRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (p *ProductRepo) Add(ctx context.Context, pr *product.Product) error {
	query, args, err := database.StatementBuilder().
		Insert("product").
		Columns("name", "category", "description", "brand", "cost", "currency", "created", "updated").
		Values(pr.Name, pr.Category, pr.Description, pr.Brand, pr.Cost, pr.Currency, pr.Created, pr.Updated).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return errors.Wrap(err, "Create Query error")
	}

	var id uint64
	if err := p.db.GetContext(ctx, &id, query, args...); err != nil {
		return errors.Wrap(err, "db.GetContext()")
	}

	pr.ID = id

	return nil
}

func (p *ProductRepo) Get(ctx context.Context, productID uint64) (product.Product, error) {
	query, args, err := database.StatementBuilder().
		Select("name", "category", "description", "brand", "cost", "currency", "created", "updated").
		From("product").
		Where(sq.Eq{"id": productID}).
		ToSql()

	if err != nil {
		return product.Product{}, errors.Wrap(err, "Create query error")
	}

	var pr product.Product
	if err := p.db.QueryRowxContext(ctx, query, args...).StructScan(&pr); err != nil {
		return product.Product{}, errors.Wrap(err, "sqlx.QueryRowxContext()")
	}

	return pr, nil
}

func (p *ProductRepo) List(ctx context.Context) ([]product.Product, error) {
	query, args, err := database.StatementBuilder().
		Select("name", "category", "description", "brand", "cost", "currency", "created", "updated").
		From("product").
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "Create query error")
	}

	if rows, err := p.db.QueryContext(ctx, query, args...); err == nil {
		var res []product.Product
		for rows.Next() {
			var p product.Product
			if err := rows.Scan(&p); err != nil {
				return nil, err
			}

			res = append(res, p)
		}

		return res, nil
	} else {
		return nil, err
	}
}

func (p *ProductRepo) Remove(ctx context.Context, productID uint64) error {
	query, args, err := database.StatementBuilder().
		Delete("product").
		Where(sq.Eq{"id": productID}).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "Create SQL query")
	}

	res, err := p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "db.ExecContext()")
	}

	count, err := res.RowsAffected()
	if count == 0 {
		return product.ProductNotFound
	}

	return nil
}
