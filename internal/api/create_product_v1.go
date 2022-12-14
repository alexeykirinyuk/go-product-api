package api

import (
	"context"
	"github.com/alexeykirinyuk/go-product-api/internal/service/product/dto"
	vo "github.com/alexeykirinyuk/go-product-api/internal/service/value_objects"
	pb "github.com/alexeykirinyuk/go-product-api/pkg/go-product-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *ProductAPI) CreateProductV1(
	ctx context.Context,
	req *pb.CreateProductV1Request,
) (*pb.CreateProductV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Warn().
			Err(err).
			Interface("Req", req).
			Msg("CreateProductV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	command := dto.CreateProductCommand{
		Name:        req.GetName(),
		Category:    req.GetCategory(),
		Description: req.GetDescription(),
		Brand:       req.GetBrand(),
		Cost:        req.GetCost(),
		Currency:    vo.Currency(req.GetCurrency()),
	}

	prod, err := p.s.CreateProduct(ctx, command)
	if err != nil {
		log.Warn().
			Err(err).
			Interface("Req", req).
			Msg("CreateProductV1 -- failed")

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	log.Debug().
		Uint64("ProductID", prod.ID).
		Interface("Req", req).
		Msg("CreateProductV1 -- success")

	return &pb.CreateProductV1Response{
		Product: toProtobufProduct(prod),
	}, nil
}
