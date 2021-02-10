package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

const dirConfig = ".konector"
const fileConfig = "config.json"

// Save Create config file
func (config *UserConfig) Save() error {

	homeDir := os.Getenv("HOME")

	if homeDir == "" {
		return fmt.Errorf("HOME environment variable is not defined")
	}

	if fileInfo, err := os.Stat(path.Join(homeDir, dirConfig)); os.IsNotExist(err) {
		err = os.Mkdir(path.Join(homeDir, dirConfig), 0755)

		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else if !fileInfo.IsDir() {
		return fmt.Errorf(".konector is not a directory")
	}

	file, err := os.OpenFile(path.Join(homeDir, dirConfig, fileConfig), os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		return err
	}

	defer file.Close()
	err = json.NewEncoder(file).Encode(config)

	if err != nil {
		return err
	}

	return nil
}
