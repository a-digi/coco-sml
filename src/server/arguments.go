package server

import (
	"flag"
)

// Arguments holds parsed command-line arguments
type Arguments struct {
	DataDir   string
	Config    string
	Port      int
	Action    string // "start" or "stop"
}

// ParseArguments parses CLI arguments and returns an Arguments struct
func ParseArguments() *Arguments {
	dataDir := flag.String("data-dir", "./data", "Path to data directory")
	config := flag.String("config", "config.json", "Path to config file")
	port := flag.Int("port", 2030, "Port to listen on (overrides config)")
	flag.Parse()

	action := "start"
	if flag.NArg() > 0 {
		switch flag.Arg(0) {
		case "start":
			action = "start"
		case "stop":
			action = "stop"
		}
	}

	return &Arguments{
		DataDir: *dataDir,
		Config:  *config,
		Port:    *port,
		Action:  action,
	}
}

