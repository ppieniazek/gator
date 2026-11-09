package config

import (
	"encoding/json/v2"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(homeDir, configFileName)
	if _, err := os.Stat(path); err != nil {
		return "", err
	}
	return path, nil
}

func Read() (Config, error) {
	cfgPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	cfgFile, err := os.ReadFile(cfgPath)
	if err != nil {
		return Config{}, err
	}
	var cfg Config
	if err = json.Unmarshal(cfgFile, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(*cfg)
}

func write(cfg Config) error {
	cfgPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	cfgData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(cfgPath, cfgData, 0600)
}
