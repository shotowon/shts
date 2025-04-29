package shts

import (
	"fmt"
	"os"

	"github.com/shotowon/shts/internal/shts/crypto"
)

func DecryptFromFiles(mkeyPath string, ciphertextPath string) (string, error) {
	cipherText, err := os.ReadFile(ciphertextPath)
	if err != nil {
		return "", fmt.Errorf("failed to read password file: %w", err)
	}

	masterKey, err := os.ReadFile(mkeyPath)
	if err != nil {
		return "", fmt.Errorf("failed to read master-key file: %w", err)
	}

	password, err := crypto.Decrypt(string(masterKey), string(cipherText))
	if err != nil {
		return "", fmt.Errorf("failed to decrypt password: %w", err)
	}

	return password, nil
}
