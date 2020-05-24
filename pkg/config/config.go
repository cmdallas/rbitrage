package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config struct for rbitrage configuration
type Config struct {
	Providers []string `yaml:"providers"`
}

// NewConfig returns a new decoded Config
func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

// ValidateConfigPath ensures the path provided is a file
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory", path)
	}

	return nil
}
