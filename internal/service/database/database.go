package database

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

var st = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func StatementBuilder() sq.StatementBuilderType {
	return st
}

// New returns DB
func New(ctx context.Context, dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create database connection")

		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "db.PingContext()")
	}

	return db, nil
}
