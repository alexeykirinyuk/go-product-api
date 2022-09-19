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

func (p *ProductAPI) RemoveProductV1(
	ctx context.Context,
	req *pb.RemoveProductV1Request,
) (*pb.RemoveProductV1Response, error) {
	prod, err := p.s.DeleteProduct(ctx, req.GetProductId())
	if errors.Is(err, product.ProductNotFound) {
		log.Debug().
			Err(err).
			Uint64("ProductID", req.GetProductId()).
			Interface("Req", req).
			Msg("RemoveProductV1 - products not found")

		totalProductNotFound.Inc()

		return nil, status.Error(codes.NotFound, "products not found")
	}

	if err != nil {
		log.Error().
			Err(err).
			Uint64("ProductID", req.GetProductId()).
			Interface("Req", req).
			Msg("RemoveProductV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().
		Uint64("ProductID", req.GetProductId()).
		Interface("Req", req).
		Msg("RemoveProductV1 -- success")

	return &pb.RemoveProductV1Response{Product: toProtobufProduct(prod)}, nil
}
