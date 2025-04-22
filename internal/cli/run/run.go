package run

import (
	"fmt"

	"github.com/spf13/cobra"
)

type runConfig struct {
	ConfigPath string
	MasterKey  string
	MasterKeys []string
}

var (
	runCfg = runConfig{}
)

var Command = &cobra.Command{
	Use:   "run",
	Short: "This command runs multiple sshuttles",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(runCfg.ConfigPath)
		fmt.Println(runCfg.MasterKey)
		fmt.Println(runCfg.MasterKeys)
		return nil
	},
}

func init() {
	Command.PersistentFlags().StringSliceVarP(&runCfg.MasterKeys, "keys", "K", nil, "Pass master key file for each sshuttle (must be listed in order in which sshutles are listed in config file)")
	Command.PersistentFlags().StringVarP(&runCfg.MasterKey, "key", "k", "", "This single master key file will be used to connect to all shhuttles in configuration file")
	Command.PersistentFlags().StringVarP(&runCfg.ConfigPath, "config", "c", "", "YAML configuration file (list sshuttles here)")
}
