package password

import (
	"fmt"
	"os"

	"github.com/shotowon/shts/internal/shts/crypto"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

type newConfig struct {
	MasterKeyFile string
	Name          string
}

var newCfg = newConfig{}

var newCmd = &cobra.Command{
	Use:  "new",
	Args: cobra.NoArgs,
	RunE: encryptRun,
}

func encryptRun(cmd *cobra.Command, args []string) error {
	if len(newCfg.MasterKeyFile) == 0 {
		return fmt.Errorf("master key not provided")
	}

	cmd.Println("password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return fmt.Errorf("failed to read password")
	}

	masterKey, err := os.ReadFile(newCfg.MasterKeyFile)
	if err != nil {
		return fmt.Errorf("failed to read master-key file: %w", err)
	}

	encrypted, err := crypto.Encrypt(string(masterKey), string(password))
	if err != nil {
		return fmt.Errorf("failed to encrypt password: %w", err)
	}

	if len(newCfg.Name) != 0 {
		file, err := os.OpenFile(newCfg.Name, os.O_WRONLY|os.O_CREATE, 0644)
		if os.IsNotExist(err) {
			return fmt.Errorf("out file does not exist: %w", err)
		}
		if os.IsPermission(err) {
			return fmt.Errorf("out file permission error: %w", err)
		}
		if err != nil {
			return fmt.Errorf("out file error: %w", err)
		}

		defer file.Close()

		if _, err := file.Write([]byte(encrypted)); err != nil {
			return fmt.Errorf("failed to write password to file: %w", err)
		}
		return nil
	}

	cmd.Println(encrypted)
	return nil
}

func init() {
	newCmd.Flags().StringVarP(&newCfg.MasterKeyFile, "key", "k", "", "master key used to sign password")
	newCmd.Flags().StringVarP(&newCfg.Name, "name", "n", "", "filename where encrypted password will be written")
}
