package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type InstallPaths struct {
	Paths []string `json:"paths"`
}

type Config struct {
	installPaths InstallPaths
}

func New() (*Config, error) {
	configDir, err := getConfigDir()

	if err != nil {
		return &Config{}, fmt.Errorf("Failed to get config directory path: %w", err)
	}

	err = os.MkdirAll(configDir, 0755)

	if err != nil {
		return &Config{}, fmt.Errorf("Failed to create config directory path: %w", err)
	}

	return &Config{}, nil
}

func getConfigDir() (string, error) {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return "", err
	}

	return path.Join(configDir, "Raijinsoft", "addoninfo"), nil
}

func getConfigFileName() (string, error) {
	configDir, err := getConfigDir()

	if err != nil {
		return "", err
	}

	return path.Join(configDir, "config.json"), nil
}

func (config *Config) Load() error {
	configDir, err := getConfigFileName()

	if err != nil {
		return err
	}

	file, err := os.Open(configDir)

	if err != nil {
		return err
	}

	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&config.installPaths)

	if err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return nil
}

func (config *Config) Save() error {
	configDir, err := getConfigFileName()

	if err != nil {
		return err
	}

	file, err := os.Create(configDir)

	if err != nil {
		return err
	}

	jsonParser := json.NewEncoder(file)
	jsonParser.SetIndent("", "    ")

	if err := jsonParser.Encode(config.installPaths); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return nil
}

func (config *Config) AddInstallPath(path string) {
	config.installPaths.Paths = append(config.installPaths.Paths, path)
}

func (config *Config) GetInstallPaths() []string {
	return config.installPaths.Paths
}
