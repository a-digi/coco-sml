package main

import (
	"fmt"
	"os"
	"github.com/a-digi/coco-sml/src/server"
)

func main() {
	args := server.ParseArguments()
	fmt.Println("coco-sml semantic search API server (starter)")

	if args.Action == "stop" {
		server.StopServer(args.DataDir)
		return
	}

	// Only if action is "start", further arguments are evaluated
	if args.Action != "start" {
		fmt.Println("No valid action provided. Use 'start' or 'stop'.")
		return
	}

	configPath := args.Config
	cfg, err := server.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}
	// Override port if provided and not zero
	if args.Port != 0 {
		cfg.Port = args.Port
	}

	// Override data directory only if argument is set and not empty
	if args.DataDir != "" {
		cfg.DataFolderPath = args.DataDir
	}

	// Start server in a separate goroutine so the terminal remains usable
	go server.StartServerWithConfig(cfg)

	// Main process stays active, but does not block the terminal
	select {}
}