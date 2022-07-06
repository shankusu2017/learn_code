package main

/* Echo client
 */

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	port := os.Args[1]

	addr := "127.0.0.1:" + port

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "connect %s fail!\n", addr)
		return
	} else {
		fmt.Fprintf(os.Stdout, "connect done!\n")
	}
	defer conn.Close()

	var buf [512]byte

	n, err := conn.Write([]byte("test client"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write socket fail! %s\n", err.Error())
		return
	}
	if n != 0 {
		// send done
	}

	n, err = conn.Read(buf[0:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "read socket fail! %s \n", err.Error())
		return
	}

	s := string(buf[0:n])
	fmt.Fprintf(os.Stdout, "recv from server: %s", s)
}
