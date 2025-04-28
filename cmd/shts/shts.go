package main

import (
	"github.com/shotowon/shts/internal/cli"
)

func main() {
	/*
		key := "AAEEBBCCDDFFEE11"
		msg := "hehedsadas"

		encrypted, err := crypto.Encrypt(key, msg)
		if err != nil {
			panic(err)
		}

		fmt.Printf("encrypted: %s\n", encrypted)

		decrypted, err := crypto.Decrypt(key, encrypted)
		if err != nil {
			panic(err)
		}

		fmt.Printf("decrypted: %s\n", decrypted)

		os.Exit(0)
	*/
	cli.Execute()
}
