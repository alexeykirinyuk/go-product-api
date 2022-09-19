package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/alexeykirinyuk/go-product-api/internal/service/database"
	"github.com/alexeykirinyuk/go-product-api/internal/service/product"
	"github.com/alexeykirinyuk/go-product-api/internal/service/product/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (p *ProductRepo) AddProduct(ctx context.Context, pr *model.Product, tx *sqlx.Tx) error {
	query, args, err := database.StatementBuilder().
		Insert("product").
		Columns("name", "category", "description", "brand", "cost", "currency", "created", "updated").
		Values(pr.Name, pr.Category, pr.Description, pr.Brand, pr.Cost, pr.Currency, pr.Created, pr.Updated).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return errors.Wrap(err, "Create Query error")
	}

	var getContexter database.GetContexter
	if tx == nil {
		getContexter = p.db
	} else {
		getContexter = tx
	}

	var id uint64
	if err := getContexter.GetContext(ctx, &id, query, args...); err != nil {
		return errors.Wrap(err, "db.GetContext()")
	}

	pr.ID = id

	return nil
}

func (p *ProductRepo) GetProduct(ctx context.Context, productID uint64, tx *sqlx.Tx) (model.Product, error) {
	query, args, err := database.StatementBuilder().
		Select("id", "name", "category", "description", "brand", "cost", "currency", "created", "updated").
		From("product").
		Where(sq.Eq{"id": productID}).
		ToSql()

	if err != nil {
		return model.Product{}, errors.Wrap(err, "Create query error")
	}

	var queryer sqlx.QueryerContext
	if tx == nil {
		queryer = p.db
	} else {
		queryer = tx
	}

	var pr model.Product
	if err := queryer.QueryRowxContext(ctx, query, args...).StructScan(&pr); err != nil {
		return model.Product{}, errors.Wrap(err, "sqlx.QueryRowxContext()")
	}

	return pr, nil
}

func (p *ProductRepo) ListProducts(ctx context.Context) ([]model.Product, error) {
	query, args, err := database.StatementBuilder().
		Select("id", "name", "category", "description", "brand", "cost", "currency", "created", "updated").
		From("product").
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "Create query error")
	}
	rows, err := p.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "QueryContext return err")
	}

	var res []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.StructScan(&p); err != nil {
			return nil, err
		}

		res = append(res, p)
	}

	return res, nil
}

func (p *ProductRepo) RemoveProduct(ctx context.Context, productID uint64, tx *sqlx.Tx) error {
	query, args, err := database.StatementBuilder().
		Delete("product").
		Where(sq.Eq{"id": productID}).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "Create SQL query")
	}

	var execer sqlx.ExecerContext
	if tx == nil {
		execer = p.db
	} else {
		execer = tx
	}

	res, err := execer.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "db.ExecContext()")
	}

	count, err := res.RowsAffected()
	if count == 0 {
		return product.ProductNotFound
	}

	return nil
}
