// Package main defines project entry
package main

import (
	"flag"
	"log"

	grpcserver "github.com/ezex-io/ezex-notification/api/grpc"
	"github.com/ezex-io/ezex-notification/internal/adapters/smtp"
	"github.com/ezex-io/ezex-notification/internal/config"
	"github.com/ezex-io/ezex-notification/internal/interactors"
)

func main() {
	configPath := flag.String("c", "", "Path to configuration file")

	flag.Parse()

	if *configPath == "" {
		log.Fatal("Please specify a config file using -c flag")
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("failed to load module's config: %v", err)
	}

	emailSender := smtp.NewSMTPAdapter(cfg.SMTP)
	emailWorker := interactors.NewEmailWorker(emailSender)

	// Create notification service (only needs the emailWorker)
	notificationService := grpcserver.NewNotificationService(emailWorker)

	// Create and start gRPC server
	server, err := grpcserver.NewServer(notificationService, grpcserver.Config{
		Port: cfg.GRPC.Port,
	})
	if err != nil {
		log.Fatalf("failed to create gRPC server: %v", err)
	}

	// Start the server in a goroutine
	go server.Start()

	log.Printf("Starting gRPC server on port %s", cfg.GRPC.Port)

	// Wait for server error
	if err := <-server.Notify(); err != nil {
		log.Fatalf("gRPC server error: %v", err)
	}
}
