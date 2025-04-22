package shts

import (
	"fmt"
	"os"

	"github.com/shotowon/shts/internal/cli/run"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "shts",
	Short: "Stupid program to run multiple sshuttles in one line",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	root.AddCommand(run.Command)
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
