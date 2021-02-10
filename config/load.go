package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

// Load config
func Load() (UserConfig, error) {

	var config UserConfig
	homeDir := os.Getenv("HOME")

	if homeDir == "" {
		return config, fmt.Errorf("HOME environment variable is not defined")
	}

	file, err := os.OpenFile(path.Join(homeDir, dirConfig, fileConfig), os.O_RDONLY, 0644)

	if err != nil {
		return config, err
	}

	defer file.Close()
	err = json.NewDecoder(file).Decode(&config)

	if err != nil {
		return config, err
	}

	return config, nil
}
