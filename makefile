generate: gens

gens:
	protoc \
	  --proto_path=brand-service/proto/v1 \
	  --go_out=brand-service/gen/go/brand/v1 \
	  --go_opt=paths=source_relative \
	  --go-grpc_out=brand-service/gen/go/brand/v1 \
	  --go-grpc_opt=paths=source_relative \
	  brand-service/proto/v1/brand.proto
