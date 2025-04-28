package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	v1 "github.com/shotowon/shts/internal/config/v1"
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
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Config: failed to open config file: %w", err)
	}

	cfg := new(Config)
	if err = cleanenv.ParseYAML(file, cfg); err != nil {
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
