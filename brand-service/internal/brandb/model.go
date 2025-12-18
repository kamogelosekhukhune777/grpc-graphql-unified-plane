package brandb

import brandpb "github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/gen/go/brand/v1"

type brandDB struct {
	ID string `db:"id"`
}

func toDBProduct(bus brandpb.Brand) brandDB {
	db := brandDB{
		ID: bus.Id,
	}

	return db
}
