package v1

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	v1 "github.com/shotowon/shts/internal/config/v1"
	"github.com/shotowon/shts/internal/shts"
)

func Run(cfg *v1.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var err error

	for _, conn := range cfg.Conns {
		func(conn v1.Conn) {
			go func() {
				err = Exec(ctx, &conn)

				if err != nil {
					cancel()
				}
			}()
		}(conn)
	}

	<-ctx.Done()

	if err != nil {
		return err
	}

	return nil
}

func Exec(ctx context.Context, conn *v1.Conn) error {
	var err error
	switch {
	case conn.PrivateKey != nil:
		err = execPKey(ctx, conn)
	case conn.MasterKey != nil && conn.Password != nil:
		err = execPassword(ctx, conn)
	default:
		err = fmt.Errorf("credentials for conn (%s) with subnets (%s) were not provided", conn.Remote, strings.Join(conn.Subnets, " "))
	}

	if err != nil {
		return err
	}

	return nil
}

func execPassword(ctx context.Context, conn *v1.Conn) error {
	args := make([]string, 0, len(conn.Subnets)+2)
	args = append(args, "-r")
	args = append(args, conn.Remote)
	args = append(args, conn.Subnets...)

	var sshCmd string

	switch conn.AcceptHostKey {
	case "yes", "no":
		sshCmd = fmt.Sprintf("ssh -o StrictHostKeyChecking=%s", conn.AcceptHostKey)
	default:
		return fmt.Errorf("expected 'yes' or 'no' for accept-host-key, got=%s", conn.AcceptHostKey)
	}

	args = append(args, "-e")
	args = append(args, sshCmd)

	cmd := exec.CommandContext(ctx, "sshuttle", args...)

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "SSH_ASKPASS_REQUIRE=force")
	cmd.Env = append(cmd.Env, fmt.Sprintf("SSH_ASKPASS=%s", os.Args[0]))
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", shts.EnvCommand, shts.CmdAskpass))
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", shts.EnvMkeyFile, *conn.MasterKey))
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", shts.EnvPassFile, *conn.Password))

	fmt.Printf("Connecting to %s with subnets %s\n", conn.Remote, strings.Join(conn.Subnets, " "))
	stderr, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to exec sshutle via password: %w\nout: %s", err, string(stderr))
	}

	return nil
}

func execPKey(ctx context.Context, conn *v1.Conn) error {
	args := make([]string, 0, len(conn.Subnets)+2)
	args = append(args, "-r")
	args = append(args, conn.Remote)
	args = append(args, conn.Subnets...)

	var sshCmd string

	switch conn.AcceptHostKey {
	case "yes", "no":
		sshCmd = fmt.Sprintf("ssh -o StrictHostKeyChecking=%s -i %s", conn.AcceptHostKey, *conn.PrivateKey)
	default:
		return fmt.Errorf("expected 'yes' or 'no' for accept-host-key, got=%s", conn.AcceptHostKey)
	}

	args = append(args, "-e")
	args = append(args, sshCmd)

	cmd := exec.CommandContext(ctx, "sshuttle", args...)

	fmt.Printf("Connecting to %s with subnets %s\n", conn.Remote, strings.Join(conn.Subnets, " "))
	stderr, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to exec sshutle using private key: %w\nout: %s", err, string(stderr))
	}

	return nil
}
