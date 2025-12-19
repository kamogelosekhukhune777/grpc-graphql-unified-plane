package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	brandpb "github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/gen/go/brand/v1"
	"github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/internal/core/business"
	repo "github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/internal/db/brandb"
	"github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/internal/db/migrations"
	"github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/internal/service"
	"github.com/kamogelosekhukhune777/grpc-graphql-unified-plane/brand-service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 1. Initialize Logger
	log := logger.New(os.Stdout, logger.LevelInfo, "BRAND-SERVICE", nil)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *logger.Logger) error {
	// 2. Initialize Repository (Storer)
	// Note: You'll need a concrete implementation of your Storer interface here
	// For now, we assume a 'NewRepository' function exists in your repo package.
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	if err := migrations.InitSchema(ctx, db); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	storer, err := repo.NewStore(log, db)
	if err != nil {
		return fmt.Errorf("failed to initialize repository: %w", err)
	}

	// 3. Initialize Business Logic
	biz := business.NewBusiness(log, storer)

	// 4. Initialize gRPC Service
	brandService := service.NewBrandService(log, biz)

	// 5. Setup gRPC Server
	server := grpc.NewServer()
	brandpb.RegisterBrandServiceServer(server, brandService)

	// Optional: Enable reflection for tools like Postman or Evans CLI
	reflection.Register(server)

	// 6. Start Listener
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	// 7. Graceful Shutdown handling
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Info(ctx, "startup", "status", "gRPC server started", "port", port)
		if err := server.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			log.Error(ctx, "shutdown", "error", err)
		}
	}()

	// Wait for control-c or kill signal
	sig := <-shutdown
	log.Info(ctx, "shutdown", "status", "shutdown started", "signal", sig.String())

	// Gracefully stop the server
	server.GracefulStop()
	log.Info(ctx, "shutdown", "status", "shutdown complete")

	return nil
}
