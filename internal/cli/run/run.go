package run

import (
	"errors"
	"fmt"
	"strings"

	"github.com/shotowon/shts/internal/config"
	"github.com/shotowon/shts/internal/shts/sshuttle"
	"github.com/spf13/cobra"
)

type commandConfig struct {
	ConfigPath string
}

var (
	commandCfg = commandConfig{}
)

var Command = &cobra.Command{
	Use:   "run",
	Short: "This command runs multiple sshuttles",
	RunE:  commandRunE,
}

func commandRunE(cmd *cobra.Command, args []string) error {
	if strings.TrimSpace(commandCfg.ConfigPath) == "" {
		return errors.New("run: path to config must be specified")
	}

	cfg, err := config.Parse(commandCfg.ConfigPath)
	if err != nil {
		return fmt.Errorf("run: failed to parse config: %w", err)
	}

	if err = sshuttle.Run(cfg); err != nil {
		return fmt.Errorf("run: failed to exec sshuttles")
	}

	return nil
}

func init() {
	Command.PersistentFlags().StringVarP(&commandCfg.ConfigPath, "config", "c", "", "YAML configuration file (list sshuttles here)")
}
