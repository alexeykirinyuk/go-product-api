package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/alexeykirinyuk/go-product-api/internal/service/database"
	"github.com/alexeykirinyuk/go-product-api/internal/service/database/migrations"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"reflect"

	"github.com/alexeykirinyuk/go-product-api/internal/config"
	"github.com/alexeykirinyuk/go-product-api/internal/server"
	"github.com/alexeykirinyuk/go-product-api/internal/tracer"
)

var (
	batchSize uint = 2
)

type Person struct {
}

func main() {
	per := int64(0)
	sl2 := &per
	sl3 := &sl2
	sl4 := &sl3

	fmt.Printf("type == %v\n", reflect.TypeOf(sl4))

	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	migration := flag.Bool("migration", false, "Defines the migration start option")
	flag.Parse()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	if cfg.Project.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	ctx := context.Background()
	db, err := database.New(ctx, dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init postgres")
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error().
				Err(err).
				Msg("Failed close db")
		}
	}()

	if *migration {
		goose.SetBaseFS(migrations.EmbedFS)
		if err = goose.Up(db.DB, "."); err != nil {
			log.Error().Err(err).Msg("Migration failed")

			return
		}

		log.Info().Msg("Migration success")
		return
	}

	tracing, err := tracer.NewTracer(&cfg)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed init tracing")

		return
	}

	defer func() {
		err := tracing.Close()

		log.Error().
			Err(err).
			Msg("Failed close tracing")
	}()

	if err := server.NewGrpcServer(db, batchSize).Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")

		return
	}
}
