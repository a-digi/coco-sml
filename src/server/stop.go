package server

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

// StopServer stops the running server gracefully using the PID file
func StopServer(dataFolderPath string) {
	pidFile := filepath.Join(dataFolderPath, "server.pid")
	data, err := os.ReadFile(pidFile)

	if err != nil {
		fmt.Println("No running server found (PID file missing)")
		return
	}

	pid, err := strconv.Atoi(string(data))
	if err != nil {
		fmt.Printf("Invalid PID in file: %v\n", err)
		return
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("Process with PID %d not found\n", pid)
		return
	}
	// Try to send SIGTERM for graceful shutdown
	if err := process.Signal(syscall.SIGTERM); err != nil {
		fmt.Printf("Failed to stop process %d: %v\n", pid, err)
		return
	}
	fmt.Printf("Sent SIGTERM to server process %d\n", pid)
}

