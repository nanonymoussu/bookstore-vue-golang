package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nanonymoussu/bookstore/internal/config"
	"github.com/nanonymoussu/bookstore/internal/db"
	"github.com/nanonymoussu/bookstore/internal/handler"
	"github.com/nanonymoussu/bookstore/internal/models"
	"github.com/nanonymoussu/bookstore/internal/repository"
	"github.com/nanonymoussu/bookstore/internal/service"
	pb "github.com/nanonymoussu/bookstore/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	database, err := db.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate database
	if err := models.AutoMigrate(database); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Create repository
	bookRepo := repository.NewBookRepository(database)

	// Create service
	bookService := service.NewBookService(bookRepo)

	// Create gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	bookstoreServer := handler.NewBookstoreServer(bookService)
	pb.RegisterBookstoreServiceServer(server, bookstoreServer)

	// Enable reflection for tools like grpcurl
	reflection.Register(server)

	log.Printf("Server listening on port %d", cfg.Server.Port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
