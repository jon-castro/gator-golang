package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(currentUserName string) error {
	c.CurrentUserName = currentUserName

	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	var readConfigFile Config
	jsonFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(data, &readConfigFile)
	if err != nil {
		return Config{}, err
	}

	return readConfigFile, nil
}

func getConfigFilePath() (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	jsonFilePath := filepath.Join(userHome, configFileName)

	return jsonFilePath, nil
}

func write(cfg Config) error {
	jsonFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
