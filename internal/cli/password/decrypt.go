package password

import (
	"fmt"
	"os"

	"github.com/shotowon/shts/internal/shts"
	"github.com/spf13/cobra"
)

type decryptConfig struct {
	MasterKeyFile  string
	CipherTextFile string
	Out            string
}

var decryptCfg = decryptConfig{}

var decrypt = &cobra.Command{
	Use:  "decrypt",
	Args: cobra.NoArgs,
	RunE: decryptRun,
}

func decryptRun(cmd *cobra.Command, args []string) error {
	if len(decryptCfg.MasterKeyFile) == 0 {
		return fmt.Errorf("empty master-key file flag")
	}
	if len(decryptCfg.CipherTextFile) == 0 {
		return fmt.Errorf("empty cipher-text file flag")
	}

	password, err := shts.DecryptFromFiles(decryptCfg.MasterKeyFile, decryptCfg.CipherTextFile)
	if err != nil {
		return err
	}

	if len(decryptCfg.Out) != 0 {
		file, err := os.OpenFile(decryptCfg.Out, os.O_WRONLY|os.O_CREATE, 0644)
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

		if _, err := file.Write([]byte(password)); err != nil {
			return fmt.Errorf("failed to write ciphertext: %w", err)
		}

		return nil
	}

	cmd.Println(password)
	return nil
}

func init() {
	decrypt.Flags().StringVarP(&decryptCfg.CipherTextFile, "cipher-text", "c", "", "ciphertext file with Encrypted password")
	decrypt.Flags().StringVarP(&decryptCfg.MasterKeyFile, "key", "k", "", "master key used to sign password")
	decrypt.Flags().StringVarP(&decryptCfg.Out, "out", "o", "", "Output file where password will be written")
}
