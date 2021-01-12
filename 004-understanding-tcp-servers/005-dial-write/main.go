package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(conn)
	}
	defer conn.Close()
	for i := 0; i < 10; i++ {
		fmt.Fprintln(conn, "dial you", i+1)
	}
}
