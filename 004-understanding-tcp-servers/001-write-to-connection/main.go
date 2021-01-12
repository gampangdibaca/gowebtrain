package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		io.WriteString(conn, "\nHello From the other side\n")
		fmt.Fprintln(conn, "How are you?")
		fmt.Fprintf(conn, "%v", "Well, i hope!")
		conn.Close()
	}
}
