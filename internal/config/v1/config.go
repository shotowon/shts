package v1

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Conns []Conn `yaml:"connections"`
}

type Conn struct {
	Remote     string   `yaml:"remote"`
	Subnets    []string `yaml:"subnets"`
	PrivateKey *string  `yaml:"private-key,omitempty"`
	Password   *string  `yaml:"password,omitempty"`
	MasterKey  *string  `yaml:"master-key,omitempty"`
}

func Parse(filepath string) (*Config, error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("V1: failed to open config file: %w", err)
	}

	cfg := new(Config)
	if err = cleanenv.ParseYAML(file, cfg); err != nil {
		return nil, fmt.Errorf("V1: failed to parse YAML of config file: %w", err)
	}
	return cfg, nil
}
