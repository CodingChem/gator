package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DB_CON_STRING string `json:"db_url"`
	CurrentUser   string `json:"current_user_name"`
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}
	var conf Config
	err = json.Unmarshal(fileData, &conf)
	if err != nil {
		return Config{}, err
	}
	return conf, nil
}

func (c *Config) SetUser(userName string) error {
	if !validUserName(userName) {
		return fmt.Errorf("Error: %s not valid username!", userName)
	}
	c.CurrentUser = userName
	err := write(c)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (path string, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path = home + string(os.PathSeparator) + configFileName
	return path, nil
}

func write(conf *Config) error {
	data, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func validUserName(userName string) bool {
	return len(userName) > 2
}
