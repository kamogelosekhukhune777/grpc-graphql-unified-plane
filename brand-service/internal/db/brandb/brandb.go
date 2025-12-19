package brandb

import (
	"context"
	"database/sql"

	brandpb "github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/gen/go/brand/v1"
	"github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/pkg/logger"
)

type Store struct {
	log *logger.Logger
	DB  *sql.DB
}

func NewStore(log *logger.Logger, db *sql.DB) (*Store, error) {
	return &Store{log: log, DB: db}, nil
}

func (s *Store) CreateBrand(ctx context.Context, brand *brandpb.Brand) (*brandpb.Brand, error) {
	_, err := s.DB.ExecContext(
		ctx, "INSERT INTO brands (id,name,description,created_at,updated_at) VALUES ($1,$2,$3,$4,$5)",
		brand.Id, brand.Name, brand.Description, brand.CreatedAt, brand.UpdatedAt)
	if err != nil {
		s.log.Error(ctx, "failed to create brand", err)
		return nil, err
	}
	s.log.Info(ctx, "created brand", brand)

	return brand, nil
}

func (s *Store) DeleteBrand(ctx context.Context, id string) error {
	_, err := s.DB.ExecContext(ctx, "DELETE FROM brands WHERE id = $1", id)
	if err != nil {
		s.log.Error(ctx, "failed to delete brand", err)
		return err
	}
	s.log.Info(ctx, "deleted brand", id)

	return nil
}

func (s *Store) UpdateBrand(ctx context.Context, brand *brandpb.Brand) (*brandpb.Brand, error) {
	_, err := s.DB.ExecContext(ctx,
		"UPDATE brands SET name = $1, description = $2, updated_at = $3 WHERE id = $4",
		brand.Name, brand.Description, brand.UpdatedAt, brand.Id)
	if err != nil {
		s.log.Error(ctx, "failed to update brand", err)
		return nil, err
	}
	s.log.Info(ctx, "updated brand", brand)

	return brand, nil
}

func (s *Store) GetBrand(ctx context.Context, id string) (*brandpb.Brand, error) {
	row := s.DB.QueryRowContext(ctx, "SELECT id, name, description, created_at, updated_at FROM brands WHERE id = $1", id)
	var brand brandpb.Brand

	err := row.Scan(&brand.Id, &brand.Name, &brand.Description, &brand.CreatedAt, &brand.UpdatedAt)
	if err != nil {
		s.log.Error(ctx, "failed to get brand", err)
		return nil, err
	}
	s.log.Info(ctx, "got brand", &brand)

	return &brand, nil
}

func (s *Store) ListBrands(ctx context.Context, pageToken string, pageSize int) ([]*brandpb.Brand, error) {
	rows, err := s.DB.QueryContext(ctx, "SELECT id, name, description, created_at, updated_at FROM brands ORDER BY id ASC LIMIT $1 OFFSET $2", pageSize, pageSize)
	if err != nil {
		s.log.Error(ctx, "failed to query brands", err)
		return nil, err
	}
	defer rows.Close()

	var brands []*brandpb.Brand
	for rows.Next() {
		var brand brandpb.Brand
		err := rows.Scan(&brand.Id, &brand.Name, &brand.Description, &brand.CreatedAt, &brand.UpdatedAt)
		if err != nil {
			s.log.Error(ctx, "failed to scan brand", err)
			return nil, err
		}
		brands = append(brands, &brand)
	}
	s.log.Info(ctx, "got all brands", len(brands))

	return brands, nil
}

func (s *Store) Close() error {
	return s.DB.Close()
}

func (s *Store) Ping() error {
	return s.DB.Ping()
}
