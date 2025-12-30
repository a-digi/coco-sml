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

	// Nur wenn Action "start" ist, werden weitere Argumente ausgewertet
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
	// Port überschreiben, falls per Argument gesetzt und ungleich 0
	if args.Port != 0 {
		cfg.Port = args.Port
	}

	// DataDir nur überschreiben, wenn Argument gesetzt und nicht leer
	if args.DataDir != "" {
		cfg.DataFolderPath = args.DataDir
	}

	server.StartServerWithConfig(cfg)
}