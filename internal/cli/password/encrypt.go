package password

import (
	"fmt"
	"os"

	"github.com/shotowon/shts/internal/shts/crypto"
	"github.com/spf13/cobra"
)

type encryptConfig struct {
	MasterKeyFile  string
	CipherTextFile string
	Out            string
}

var encryptCfg = encryptConfig{}

var encrypt = &cobra.Command{
	Use:  "encrypt",
	Args: cobra.NoArgs,
	RunE: encryptRun,
}

func encryptRun(cmd *cobra.Command, args []string) error {
	if len(encryptCfg.MasterKeyFile) == 0 {
		return fmt.Errorf("empty master-key file flag")
	}
	if len(encryptCfg.CipherTextFile) == 0 {
		return fmt.Errorf("empty password-file flag")
	}

	password, err := os.ReadFile(encryptCfg.CipherTextFile)
	if err != nil {
		return fmt.Errorf("failed to read password file: %w", err)
	}

	masterKey, err := os.ReadFile(encryptCfg.MasterKeyFile)
	if err != nil {
		return fmt.Errorf("failed to read master-key file: %w", err)
	}

	encrypted, err := crypto.Encrypt(string(masterKey), string(password))
	if err != nil {
		return fmt.Errorf("failed to encrypt password: %w", err)
	}

	if len(encryptCfg.Out) != 0 {
		file, err := os.OpenFile(encryptCfg.Out, os.O_WRONLY|os.O_CREATE, 0644)
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
			return fmt.Errorf("failed to write ciphertext: %w", err)
		}
		return nil
	}

	cmd.Println(encrypted)
	return nil
}

func init() {
	encrypt.Flags().StringVarP(&encryptCfg.MasterKeyFile, "key", "k", "", "master key used to sign password")
	encrypt.Flags().StringVarP(&encryptCfg.CipherTextFile, "password", "p", "", "password file")
	encrypt.Flags().StringVarP(&encryptCfg.Out, "out", "o", "", "output file to write encrypted password")
}
