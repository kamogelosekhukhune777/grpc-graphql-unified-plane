package repository

import (
	"context"

	brandpb "github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/gen/go/brand/v1"
)

type Storer interface {
	GetBrand(ctx context.Context, id string) (*brandpb.Brand, error)
	CreateBrand(ctx context.Context, brand *brandpb.Brand) (*brandpb.Brand, error)
	UpdateBrand(ctx context.Context, id string, brand *brandpb.Brand) (*brandpb.Brand, error)
	DeleteBrand(ctx context.Context, id string) error
	ListBrands(ctx context.Context, page, pageSize int) ([]*brandpb.Brand, error)
}
