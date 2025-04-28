package cli

import (
	"fmt"
	"os"

	"github.com/shotowon/shts/internal/cli/password"
	"github.com/shotowon/shts/internal/cli/run"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "shts",
	Short: "Stupid program to run multiple sshuttles in one line",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil
	},
}

func init() {
	root.AddCommand(run.Command)
	root.AddCommand(password.Command)
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
