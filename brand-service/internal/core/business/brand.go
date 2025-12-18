// package business
package business

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	brandpb "github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/gen/go/brand/v1"
	repo "github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/internal/repository"
	"github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/pkg/logger"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Business manages the set of APIs for product access.
type Business struct {
	log    *logger.Logger
	storer repo.Storer
}

// NewBusiness creates a new business instance
func NewBusiness(log *logger.Logger, storer repo.Storer) *Business {
	return &Business{
		log:    log,
		storer: storer,
	}
}

// CreateBrand creates a new brand
func (b *Business) CreateBrand(ctx context.Context, brand *brandpb.Brand) (*brandpb.Brand, error) {
	brnd := &brandpb.Brand{
		Id:          uuid.New().String(),
		Name:        brand.Name,
		Description: brand.Description,
		CreatedAt:   timestamppb.Now(),
		UpdatedAt:   timestamppb.Now(),
	}

	ubrnd, err := b.storer.CreateBrand(ctx, brnd)
	if err != nil {
		return &brandpb.Brand{}, fmt.Errorf("create: %w", err)
	}

	return ubrnd, nil

}

// QueryBrand queries a brand by ID
func (b *Business) QueryBrandByID(ctx context.Context, id string) (*brandpb.Brand, error) {
	brnd, err := b.storer.GetBrand(ctx, id)
	if err != nil {
		return &brandpb.Brand{}, fmt.Errorf("query: brandID[%s]: %w", id, err)
	}

	return brnd, nil
}

// UpdateBrand updates a brand by ID
func (b *Business) UpdateBrand(ctx context.Context, id string, brand *brandpb.Brand) (*brandpb.Brand, error) {
	brnd, err := b.QueryBrandByID(ctx, id)
	if err != nil {
		return &brandpb.Brand{}, fmt.Errorf("from update query: brandID[%s]: %w", id, err)
	}

	brnd.Name = brand.Name
	brnd.Description = brand.Description
	brnd.UpdatedAt = timestamppb.Now()

	upbrnd, err := b.storer.UpdateBrand(ctx, brnd)
	if err != nil {
		return &brandpb.Brand{}, fmt.Errorf("update: %w", err)
	}

	return upbrnd, nil
}

// DeleteBrand deletes a brand by ID
func (b *Business) DeleteBrand(ctx context.Context, id string) error {
	err := b.storer.DeleteBrand(ctx, id)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

// ListBrands lists brands by page and page size
func (b *Business) ListBrands(ctx context.Context, page, pageSize int) ([]*brandpb.Brand, error) {
	brnds, err := b.storer.ListBrands(ctx, page, pageSize)
	if err != nil {
		return []*brandpb.Brand{}, fmt.Errorf("list: %w", err)
	}

	return brnds, nil
}
