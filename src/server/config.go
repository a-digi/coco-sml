package server

import (
	"encoding/json"
	"os"
)

// Config holds server configuration
type Config struct {
	Port           int    `json:"port"`
	DataFolderPath string `json:"dataFolderPath"`
}

// LoadConfig loads configuration from a JSON file
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		// Datei existiert nicht: Default-Werte zurückgeben
		return &Config{
			Port: 2030,
			DataFolderPath: "./data",
		}, nil
	}
	defer file.Close()
	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}
	if cfg.Port == 0 {
		cfg.Port = 2030
	}
	return &cfg, nil
}
