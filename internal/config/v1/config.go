package v1

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Conns []Conn `yaml:"connections"`
}

type Conn struct {
	Remote        string   `yaml:"remote"`
	Subnets       []string `yaml:"subnets"`
	PrivateKey    *string  `yaml:"private-key,omitempty"`
	Password      *string  `yaml:"password,omitempty"`
	MasterKey     *string  `yaml:"master-key,omitempty"`
	AcceptHostKey string   `yaml:"accept-host-key"`
}

func Parse(filepath string) (*Config, error) {
	fileContents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("V1: failed to open config file: %w", err)
	}

	cfg := new(Config)
	if err = yaml.Unmarshal(fileContents, cfg); err != nil {
		return nil, fmt.Errorf("V1: failed to parse YAML of config file: %w", err)
	}
	return cfg, nil
}
