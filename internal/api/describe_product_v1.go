package api

import (
	"context"
	"errors"
	"github.com/alexeykirinyuk/go-product-api/internal/service/product"
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
			Interface("Req", req).
			Msg("DescribeProductV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	prod, err := p.s.GetProductById(ctx, req.GetProductId())

	if errors.Is(err, product.ProductNotFound) {
		log.Debug().
			Err(err).
			Uint64("ProductID", req.GetProductId()).
			Interface("Req", req).
			Msg("DescribeProductV1 - p not found")

		totalProductNotFound.Inc()

		return nil, status.Error(codes.NotFound, "p was not found")
	}

	if err != nil {
		log.Error().
			Err(err).
			Uint64("ProductID", req.GetProductId()).
			Interface("Req", req).
			Msg("DescribeProductV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().
		Uint64("ProductID", req.GetProductId()).
		Interface("Req", req).
		Msg("CreateProductV1 -- success")

	return &pb.DescribeProductV1Response{
		Product: toProtobufProduct(prod),
	}, nil
}
