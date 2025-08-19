package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(currentUserName string) {
	c.CurrentUserName = currentUserName
}

func Read() (Config, error) {
	var readConfigFile Config
	userHome, _ := os.UserHomeDir()
	jsonFile := filepath.Join(userHome, configFileName)
	decoder := json.NewDecoder(strings.NewReader(jsonFile))
	if err := decoder.Decode(&readConfigFile); err != nil {
		return Config{}, err
	}
	return readConfigFile, nil
}
