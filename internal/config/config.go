package config

import (
	"fmt"
	"os"

	v1 "github.com/shotowon/shts/internal/config/v1"
	"gopkg.in/yaml.v3"
)

type Version int

const (
	V1 Version = 1
)

type Config struct {
	Version Version `yaml:"version"`
	Content any     `yaml:"-"`
}

func Parse(filepath string) (*Config, error) {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Config: failed to open config file: %w", err)
	}

	cfg := new(Config)
	if err = yaml.Unmarshal(fileContent, cfg); err != nil {
		return nil, fmt.Errorf("Config: failed to parse YAML of base config file: %w", err)
	}

	switch cfg.Version {
	case V1:
		v1Cfg, err := v1.Parse(filepath)
		if err != nil {
			return nil, fmt.Errorf("Config: %w", err)
		}

		cfg.Content = v1Cfg
	}

	return cfg, nil
}
