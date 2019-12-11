package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Panicln(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panicln(err)
		}

		io.WriteString(conn, fmt.Sprint("Hello world\n", time.Now(), "\n"))

		conn.Close()
	}
}
