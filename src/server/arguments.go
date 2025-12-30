package server

import (
	"flag"
)

// ActionType represents allowed actions for the CLI
// Only "start" and "stop" are valid

type ActionType int

const (
	ActionNone ActionType = iota
	ActionStart
	ActionStop
)

func (a ActionType) String() string {
	switch a {
	case ActionStart:
		return "start"
	case ActionStop:
		return "stop"
	default:
		return ""
	}
}

// Arguments holds parsed command-line arguments
type Arguments struct {
	DataDir   string
	Config    string
	Port      int
	Action    ActionType // Only ActionStart or ActionStop
}

// ParseArguments parses CLI arguments and returns an Arguments struct
func ParseArguments() *Arguments {
	dataDir := flag.String("data-dir", "./data", "Path to data directory")
	config := flag.String("config", "config.json", "Path to config file")
	port := flag.Int("port", 2030, "Port to listen on (overrides config)")
	flag.Parse()

	action := ActionNone
	if flag.NArg() > 0 {
		switch flag.Arg(0) {
		case "start":
			action = ActionStart
		case "stop":
			action = ActionStop
		}
	}

	return &Arguments{
		DataDir: *dataDir,
		Config:  *config,
		Port:    *port,
		Action:  action,
	}
}
