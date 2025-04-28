package sshuttle

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/shotowon/shts/internal/config"
	v1 "github.com/shotowon/shts/internal/config/v1"
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

func ExecPassword(masterKey string, password string, remote string, subnets []string) error {
	args := make([]string, 0, len(subnets)+2)
	args = append(args, "-r")
	args = append(args, remote)
	args = append(args, subnets...)

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("sshuttle", args...)
	cmd.Env = os.Environ()
	askpass := fmt.Sprintf("SSH_ASKPASS='%s password decrypt -p %s -k %s'", path.Join(wd, os.Args[0]), path.Join(wd, password), path.Join(wd, masterKey))
	cmd.Env = append(cmd.Env, askpass)
	cmd.Env = append(cmd.Env, "SSH_ASKPASS_REQUIRE=force")

	stderr, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to exec sshutle via password: %w\nout: %s", err, string(stderr))
	}

	return nil
}

func ExecPKey(pkey string, remote string, subnets []string) error {
	return nil
}
