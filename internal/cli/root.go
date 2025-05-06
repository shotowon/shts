package cli

import (
	"fmt"
	"os"

	"github.com/shotowon/shts/internal/cli/password"
	"github.com/shotowon/shts/internal/cli/shutle"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:          "shts",
	Short:        "Stupid program to run multiple sshuttles in one line",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			if err := cmd.Help(); err != nil {
				return err
			}
			return nil
		}
		return nil
	},
}

func init() {
	root.AddCommand(password.Command)
	root.AddCommand(shutle.Command)
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
