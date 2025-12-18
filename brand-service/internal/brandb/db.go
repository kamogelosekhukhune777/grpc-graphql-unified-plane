package brandb

import (
	"database/sql"
	"github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/pkg/logger"
)

type Store struct {
	log *logger.Logger
	DB  *sql.DB
}

func NewStore(log *logger.Logger, db *sql.DB) (*Store, error) {
	return &Store{log: log, DB: db}, nil
}
