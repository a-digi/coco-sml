package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

func isProcessRunning(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	// Signal 0 prüft nur, ob der Prozess existiert (Unix)
	err = process.Signal(syscall.Signal(0))
	return err == nil
}

func writePIDFile(pidFile string) error {
	pid := os.Getpid()
	return os.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0644)
}

func removePIDFile(pidFile string) {
	_ = os.Remove(pidFile)
}

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
		if pid, err := strconv.Atoi(string(data)); err == nil && isProcessRunning(pid) {
			log.Fatalf("Server is already running with PID %d", pid)
		}
	}
	if err := writePIDFile(pidFile); err != nil {
		log.Fatalf("Failed to write PID file: %v", err)
	}
	defer removePIDFile(pidFile)

	http.HandleFunc("/v1/status", statusHandler)
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Starting coco-sml API server on %s...", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

