package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

func Encrypt(key string, message string) (string, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to create nonce: %w", err)
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(message), nil)

	return fmt.Sprintf("%s.%s", hex.EncodeToString(ciphertext), hex.EncodeToString(nonce)), nil
}

func Decrypt(key string, encryptedMessage string) (string, error) {
	parts := strings.Split(encryptedMessage, ".")
	if len(parts) != 2 {
		return "", fmt.Errorf("failed to decrypt message: expected parts 2, got %d", len(parts))
	}

	ciphertext, err := hex.DecodeString(parts[0])
	if err != nil {
		return "", fmt.Errorf("failed to decrypt password hex: %w", err)
	}

	nonce, err := hex.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("failed to decrypt nonce hex: %w", err)
	}

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	message, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt message: %w", err)
	}

	return string(message), nil
}
