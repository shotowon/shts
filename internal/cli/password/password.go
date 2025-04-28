package password

import (
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "password",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}

		return nil
	},
}

func init() {
	Command.AddCommand(encrypt)
	Command.AddCommand(decrypt)
}
