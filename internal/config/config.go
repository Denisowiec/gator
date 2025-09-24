package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Dburl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name

	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("couldn't reconstruct json data: %v", err)
	}

	if err = os.WriteFile(filepath, data, 0644); err != nil {
		return fmt.Errorf("error writing the data back to the config file: %v", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {
	filepath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting the home directory path: %v", err)
	}
	filepath += string(os.PathSeparator) + configFileName

	return filepath, nil
}

func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, fmt.Errorf("error opening config file: %v", err)
	}
	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return Config{}, fmt.Errorf("error decoding json data: %v", err)
	}

	return config, nil
}
