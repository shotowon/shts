package password

import "github.com/spf13/cobra"

type encryptConfig struct {
	MasterKey string
}

var encryptCfg = encryptConfig{}

var Encrypt = &cobra.Command{
	Use:  "encrypt",
	Args: cobra.ExactArgs(1),
}

func init() {
	Encrypt.Flags().StringVarP(&encryptCfg.MasterKey, "key", "k", "", "master key used to sign password")
	Encrypt.Flags().StringVarP(&encryptCfg.MasterKey, "out", "o", "rename_me", "output file to write encrypted password, default (rename_me)")
}
