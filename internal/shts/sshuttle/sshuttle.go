package sshuttle

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/shotowon/shts/internal/config"
	v1 "github.com/shotowon/shts/internal/config/v1"
	"github.com/shotowon/shts/internal/shts"
)

func Run(cfg *config.Config) error {

	switch cfg.Version {
	case config.V1:

		v1Cfg, ok := cfg.Content.(*v1.Config)
		if !ok {
			return errors.New("config contents don't match config version")
		}

		return RunV1(v1Cfg)
	}

	return nil
}

func ExecPassword(mkeyPath string, passFile string, remote string, subnets []string) error {
	args := make([]string, 0, len(subnets)+2)
	args = append(args, "-r")
	args = append(args, remote)
	args = append(args, subnets...)

	cmd := exec.Command("sshuttle", args...)
	cmd.Env = os.Environ()
	askpass := fmt.Sprintf("SSH_ASKPASS='%s'", os.Args[0])
	cmd.Env = append(cmd.Env, askpass)
	cmd.Env = append(cmd.Env, "%s='%s'", shts.EnvCommand, shts.CmdAskpass)
	cmd.Env = append(cmd.Env, "%s=%s", shts.EnvMkeyFile, mkeyPath)
	cmd.Env = append(cmd.Env, "%s=%s", shts.EnvPassFile, passFile)

	stderr, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to exec sshutle via password: %w\nout: %s", err, string(stderr))
	}

	return nil
}

func ConnectPassword(mkeyPath string, passFile string, remote string) error {
	args := make([]string, 0, 1)
	args = append(args, remote)

	cmd := exec.Command("ssh", args...)
	cmd.Env = os.Environ()

	askpass := fmt.Sprintf("SSH_ASKPASS=%s", os.Args[0])
	cmd.Env = append(cmd.Env, askpass)
	cmd.Env = append(cmd.Env, "SSH_ASKPASS_REQUIRE=force")
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", shts.EnvCommand, shts.CmdAskpass))
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", shts.EnvMkeyFile, mkeyPath))
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", shts.EnvPassFile, passFile))

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to exec ssh via password: %w", err)
	}

	return nil
}

func ExecPKey(pkey string, remote string, subnets []string) error {
	return nil
}
