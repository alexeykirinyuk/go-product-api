package api

import (
	"context"
	"errors"
	"github.com/alexeykirinyuk/go-product-api/internal/shared/internal_errors"
	pb "github.com/alexeykirinyuk/go-product-api/pkg/go-product-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *ProductAPI) RemoveProductV1(
	ctx context.Context,
	req *pb.RemoveProductV1Request,
) (*pb.RemoveProductV1Response, error) {
	prod, err := p.s.DeleteProduct(ctx, req.GetProductId())
	if errors.Is(err, internal_errors.ProductNotFound) {
		log.Debug().
			Err(err).
			Uint64("ProductID", req.GetProductId()).
			Msg("RemoveProductV1 - product not found")

		totalProductNotFound.Inc()

		return nil, status.Error(codes.NotFound, "product not found")
	}

	if err != nil {
		log.Error().
			Err(err).
			Uint64("ProductID", req.GetProductId()).
			Msg("RemoveProductV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.RemoveProductV1Response{Product: toProtobufProduct(prod)}, nil
}
