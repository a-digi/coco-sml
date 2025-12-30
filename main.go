package main

import (
	"fmt"
	"github.com/a-digi/coco-sml/src/server"
	"os"
)

func main() {
	fmt.Println("coco-sml semantic search API server (starter)")
	configPath := "config.json"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}
	cfg, err := server.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}
	server.StartServerWithConfig(cfg)
}