generate-brand:
	cd ./brand-service/proto && protoc --go_out=. --go-grpc_out=. brand.proto
