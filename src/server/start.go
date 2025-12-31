package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
	"github.com/a-digi/coco-sml/src/server/process"
	"github.com/a-digi/coco-sml/src/api/config"
	"github.com/a-digi/coco-sml/src/db/binary"
	"github.com/a-digi/coco-sml/src/db/binary/orm"
    "github.com/a-digi/coco-sml/src/api/config/routes"
    "github.com/a-digi/coco-sml/src/api/config/di"
)

// StartServer starts the HTTP API server
func StartServer(addr string) {
	log.Printf("Starting coco-sml API server on %s...", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// StartServerWithConfig starts the HTTP API server with config and manages PID file
func StartServerWithConfig(cfg *Config) {
	pidFile := filepath.Join(cfg.DataFolderPath, "server.pid")

	if data, err := os.ReadFile(pidFile); err == nil {
		if pid, err := strconv.Atoi(string(data)); err == nil && process.IsProcessRunning(pid) {
			log.Fatalf("Server is already running with PID %d", pid)
		}
	}

	if err := process.WritePIDFile(pidFile); err != nil {
		log.Fatalf("Failed to write PID file: %v", err)
	}

	defer process.RemovePIDFile(pidFile)

	dbData := filepath.Join(cfg.DataFolderPath, "db")
	if err := os.MkdirAll(dbData, 0755); err != nil {
		log.Fatalf("Failed to create database folder: %v", err)
	}

	serverInstance, err := binary.Start(dbData)
	if err != nil {
		log.Fatalf("Failed to start database server: %v", err)
	}

	registryManager := orm.NewRegistryManager(serverInstance)
	serviceBag := di.NewServiceBag()
	serviceBag.SetRegistryManager(registryManager)

	rb := routes.NewRouteBuilder(serviceBag)
	config.SetupApi(serverInstance)
	routes.RegisterRoutes(rb)

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
		Handler: rb.Handler(),
	}

	done := make(chan struct{})

	// Signal handler: ignore SIGINT, use SIGTERM for shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for sig := range sigChan {
			switch sig {
			case syscall.SIGINT:
				log.Println("SIGINT (CTRL+C) ignored. Use 'make stop' or SIGTERM to stop the server.")
			case syscall.SIGTERM:
				log.Println("SIGTERM received, shutting down gracefully...")
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				if err := server.Shutdown(ctx); err != nil {
					log.Fatalf("Graceful shutdown failed: %v", err)
				}
				close(done)
				return
			}
		}
	}()

	log.Printf("Starting coco-sml API server")
    var host string
    if server.Addr[0] == ':' {
        host = "localhost" // default host if only port is set
        fmt.Printf("Server running at: http://%s%s\n", host, server.Addr)
    } else {
        fmt.Printf("Server running at: http://%s\n", server.Addr)
    }
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
	<-done // Wait for shutdown signal
}
