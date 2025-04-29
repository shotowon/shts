package main

import (
	"fmt"
	"os"

	"github.com/shotowon/shts/internal/cli"
	"github.com/shotowon/shts/internal/shts"
)

func main() {
	envCommand := os.Getenv(shts.EnvCommand)

	switch envCommand {
	case shts.CmdAskpass:
		mkeyPath := os.Getenv(shts.EnvMkeyFile)
		if len(mkeyPath) == 0 {
			fmt.Fprintln(os.Stderr, "Please provide master key to askpass")
			return
		}

		passwordPath := os.Getenv(shts.EnvPassFile)
		if len(passwordPath) == 0 {
			fmt.Fprintln(os.Stderr, "Please provide password to askpass")
			return
		}

		password, err := shts.DecryptFromFiles(mkeyPath, passwordPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		fmt.Println(password)
		return
	}

	cli.Execute()
}
