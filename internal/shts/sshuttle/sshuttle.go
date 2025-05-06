package sshuttle

import (
	"errors"

	"github.com/shotowon/shts/internal/config"
	v1cfg "github.com/shotowon/shts/internal/config/v1"
	v1 "github.com/shotowon/shts/internal/shts/sshuttle/v1"
)

func Run(cfg *config.Config) error {
	switch cfg.Version {
	case config.V1:
		v1Cfg, ok := cfg.Content.(*v1cfg.Config)
		if !ok {
			return errors.New("config contents don't match config version")
		}

		return v1.Run(v1Cfg)
	}

	return nil
}
