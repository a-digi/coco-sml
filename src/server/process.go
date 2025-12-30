package server

import (
	"os"
	"syscall"
	"strconv"
)

// isProcessRunning checks if a process with the given PID is running
func IsProcessRunning(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	// Signal 0 prüft nur, ob der Prozess existiert (Unix)
	err = process.Signal(syscall.Signal(0))
	return err == nil
}

// writePIDFile writes the current process PID to the given file
func WritePIDFile(pidFile string) error {
	pid := os.Getpid()
	return os.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0644)
}

// removePIDFile removes the PID file
func RemovePIDFile(pidFile string) {
	_ = os.Remove(pidFile)
}

