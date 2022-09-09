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
	req *pb.ListProductsV1Request,
) (*pb.ListProductsV1Response, error) {
	products, err := p.s.GetProducts(ctx)
	if err != nil {
		log.Error().
			Err(err).
			Interface("Req", req).
			Msg("ListProductsV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	res := make([]*pb.Product, 0, len(products))
	for _, p := range products {
		res = append(res, toProtobufProduct(p))
	}

	log.Debug().
		Interface("Req", req).
		Msg("ListProductsV1 -- success")

	return &pb.ListProductsV1Response{Products: res}, nil
}
