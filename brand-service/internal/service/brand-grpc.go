package service

import (
	"context"

	brandpb "github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/gen/go/brand/v1"
	"github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/internal/core/business"
	"github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/pkg/logger"
)

type BrandService struct {
	log *logger.Logger
	brandpb.UnimplementedBrandServiceServer
	business.Business
}

func NewBrandService(log *logger.Logger, business business.Business) *BrandService {
	return &BrandService{
		log:      log,
		Business: business,
	}
}

func (s *BrandService) CreateBrand(ctx context.Context, req *brandpb.CreateBrandRequest) (*brandpb.CreateBrandResponse, error) {
	brand, err := s.Business.CreateBrand(ctx, &brandpb.Brand{
		Name:        req.GetName(),
		Description: req.GetDescription(),
	})
	if err != nil {
		s.log.Error(ctx, "Failed to create brand", "error", err)
		return nil, err
	}
	s.log.Info(ctx, "Brand created successfully", "id", brand.Id)

	return &brandpb.CreateBrandResponse{
		Brand: brand,
	}, nil
}

func (s *BrandService) GetBrand(ctx context.Context, req *brandpb.GetBrandRequest) (*brandpb.GetBrandResponse, error) {
	brand, err := s.Business.QueryBrandByID(ctx, req.GetId())
	if err != nil {
		s.log.Error(ctx, "Failed to get brand", "error", err)
		return nil, err
	}
	s.log.Info(ctx, "Brand retrieved successfully", "id", brand.Id)

	return &brandpb.GetBrandResponse{
		Brand: brand,
	}, nil
}

func (s *BrandService) UpdateBrand(ctx context.Context, req *brandpb.UpdateBrandRequest) (*brandpb.UpdateBrandResponse, error) {
	brand, err := s.Business.UpdateBrand(ctx, req.Id, &brandpb.Brand{
		Id:          req.GetId(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
	})
	if err != nil {
		s.log.Error(ctx, "Failed to update brand", "error", err)
		return nil, err
	}
	s.log.Info(ctx, "Brand updated successfully", "id", req.GetId())

	return &brandpb.UpdateBrandResponse{
		Brand: brand,
	}, nil
}

func (s *BrandService) DeleteBrand(ctx context.Context, req *brandpb.DeleteBrandRequest) (*brandpb.DeleteBrandResponse, error) {
	err := s.Business.DeleteBrand(ctx, req.Id)
	if err != nil {
		s.log.Error(ctx, "Failed to delete brand", "error", err)
		return nil, err
	}
	s.log.Info(ctx, "Brand deleted successfully", "id", req.Id)

	return &brandpb.DeleteBrandResponse{}, nil
}

func (s *BrandService) ListBrands(ctx context.Context, req *brandpb.ListBrandsRequest) (*brandpb.ListBrandsResponse, error) {
	brands, err := s.Business.ListBrands(ctx, req.GetPageToken(), int(req.GetLimit()))
	if err != nil {
		s.log.Error(ctx, "Failed to list brands", "error", err)
		return nil, err
	}
	s.log.Info(ctx, "Brands listed successfully")

	return &brandpb.ListBrandsResponse{
		Brands: brands,
	}, nil
}
