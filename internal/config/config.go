package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const cfgFile = ".gatorconfig.json"

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, cfgFile), nil
}

func Read() (Config, error) {

	cfg_path, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(cfg_path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	cfg := Config{}
	config, err := os.ReadFile(cfg_path)
	json.Unmarshal(config, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func write(cfg Config) error {
	cfg_path, err := getConfigPath()
	if err != nil {
		return err
	}
	cfg_json, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	os.WriteFile(cfg_path, cfg_json, 0644)
	return nil
}

func (c *Config) SetCurrentUserName(name string) error {
	c.CurrentUserName = name
	return write(*c)
}
