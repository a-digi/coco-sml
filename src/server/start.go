package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"github.com/a-digi/coco-sml/src/server/process"
)

// StartServer starts the HTTP API server
func StartServer(addr string) {
	http.HandleFunc("/v1/status", statusHandler)
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

	http.HandleFunc("/v1/status", statusHandler)
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Starting coco-sml API server on %s...", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
