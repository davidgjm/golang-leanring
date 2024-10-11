package ch03

import (
	"fmt"
	"net"
	"os"
)

func listen() {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("error setting up listener", err)
		os.Exit(1)
	}
	fmt.Printf("Listening on port %s", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		go func(c net.Conn) {
			defer c.Close()

			fmt.Printf("Remote address: %s\n", c.RemoteAddr())
		}(conn)
	}
}
