package api

import (
	"context"
	"github.com/alexeykirinyuk/go-product-api/internal/service/dto"
	"github.com/alexeykirinyuk/go-product-api/internal/shared/enum"
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
			Msg("CreateProductV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	command := dto.CreateProductCommand{
		Name:        req.Name,
		Category:    req.Category,
		Description: req.Description,
		Brand:       req.Brand,
		Cost:        req.Cost,
		Currency:    enum.Currency(req.Currency),
	}

	prod, err := p.s.CreateProduct(ctx, command)
	if err != nil {
		log.Warn().
			Err(err).
			Msg("CreateProductV1 -- failed")

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.CreateProductV1Response{
		Product: toProtobufProduct(prod),
	}, nil
}
