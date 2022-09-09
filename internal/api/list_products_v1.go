package api

import (
	"context"
	pb "github.com/alexeykirinyuk/go-product-api/pkg/go-product-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *ProductAPI) ListProductsV1(
	ctx context.Context,
	_ *pb.ListProductsV1Request,
) (*pb.ListProductsV1Response, error) {
	products, err := p.s.GetProducts(ctx)
	if err != nil {
		log.Error().
			Err(err).
			Msg("ListProductsV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	res := make([]*pb.Product, 0, len(products))
	for _, p := range products {
		res = append(res, toProtobufProduct(p))
	}

	return &pb.ListProductsV1Response{Products: res}, nil
}
