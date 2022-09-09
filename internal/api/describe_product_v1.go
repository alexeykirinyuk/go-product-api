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

func (p *ProductAPI) DescribeProductV1(
	ctx context.Context,
	req *pb.DescribeProductV1Request,
) (*pb.DescribeProductV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().
			Err(err).
			Msg("DescribeProductV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	product, err := p.s.GetProductById(ctx, req.ProductId)

	if errors.Is(err, internal_errors.ProductNotFound) {
		log.Debug().
			Err(err).
			Uint64("ProductID", req.ProductId).
			Msg("DescribeProductV1 - product not found")

		totalProductNotFound.Inc()

		return nil, status.Error(codes.NotFound, "product was not found")
	}

	if err != nil {
		log.Error().
			Err(err).
			Uint64("ProductID", req.ProductId).
			Msg("DescribeProductV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DescribeProductV1Response{
		Product: toProtobufProduct(product),
	}, nil
}
