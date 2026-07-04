package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbUrl            string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	var cfg Config

	configPath, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName

	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(cfg); err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	userHomeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return filepath.Join(userHomeDir, ".gatorconfig.json"), nil
}
