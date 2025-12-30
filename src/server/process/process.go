package process

import (
	"os"
	"syscall"
	"strconv"
)

// IsProcessRunning checks if a process with the given PID is running
func IsProcessRunning(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	return err == nil
}

// WritePIDFile writes the current process PID to the given file
func WritePIDFile(pidFile string) error {
	pid := os.Getpid()
	return os.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0644)
}

// RemovePIDFile removes the PID file
func RemovePIDFile(pidFile string) {
	_ = os.Remove(pidFile)
}

