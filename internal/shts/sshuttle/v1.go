package sshuttle

import (
	"fmt"
	"os"
	"sync"

	v1 "github.com/shotowon/shts/internal/config/v1"
)

func RunV1(cfg *v1.Config) error {
	wg := sync.WaitGroup{}
	wg.Add(len(cfg.Conns))

	sshuttleFunc := func(conn v1.Conn) {
		if conn.MasterKey != nil && conn.Password != nil {
			err := ExecPassword(*conn.MasterKey, *conn.Password, conn.Remote, conn.Subnets)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error sshuttle: %v\n", err)
			}
			fmt.Println("donee")
			wg.Done()
		}
	}

	for _, conn := range cfg.Conns {
		go sshuttleFunc(conn)
	}
	wg.Wait()

	return nil
}
