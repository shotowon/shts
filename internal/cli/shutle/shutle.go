package shutle

import (
	"errors"
	"fmt"

	"github.com/shotowon/shts/internal/config"
	"github.com/shotowon/shts/internal/shts/sshuttle"
	"github.com/spf13/cobra"
)

type shutleConfig struct {
	cfg string
}

var cfg = shutleConfig{}

var Command = &cobra.Command{
	Use:  "shutle",
	RunE: runE,
}

func runE(cmd *cobra.Command, args []string) error {
	if len(cfg.cfg) == 0 {
		return errors.New("config not provided")
	}

	cfg, err := config.Parse(cfg.cfg)
	if err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	err = sshuttle.Run(cfg)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to sshuttle: %w", err)
	}

	return nil
}

func init() {
	Command.Flags().StringVarP(&cfg.cfg, "config", "c", "", "pass config with multiple sshuttles defined, to run multiple sshuttles at the same time")
}
