package main

import (
	"fmt"
	"net"

	"github.com/swxctx/pot/plog"
)

// newServer
func newServer() error {
	// tcp listen
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		return fmt.Errorf("port listen err-> %v", err)
	}

	defer listener.Close()

	plog.Infof("pot server is starting...")

	for {
		// listen and accept
		conn, err := listener.Accept()
		if err != nil {
			plog.Errorf("listener accept err-> %v", err)
			continue
		}
		go cmdHandleConnection(conn)
	}
}
