package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Application returns a new rbitrage application configuration.
// An Application should be thought of as a specfic user workload.
type Application struct {
	// Name The name of the specific application
	Name string `yaml:"name"`
	// Properties Properties specific to the application
	Properties struct {
		// Providers Cloud/platform providers that rbitrage will attempt to use
		Providers []struct {
			// Name The name of the cloud provider. Valid choices are ["aws", "gcp", "az"]
			Name  string `yaml:"name"`
			Nodes struct {
				TypeOverrides []string `yaml:"typeOverrides"`
				VCPU          int      `yaml:"vcpu"`
				Memory        int      `yaml:"memory"`
				GroupName     string   `yaml:"groupName"`
				MinSize       int      `yaml:"minSize"`
				MaxSize       int      `yaml:"maxSize"`
				Region        string   `yaml:"region"`
			} `yaml:"nodes"`
		} `yaml:"providers"`
	} `yaml:"properties"`
}

// Config struct for rbitrage configuration
type Config struct {
	Applications []Application `yaml:"applications"`
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
