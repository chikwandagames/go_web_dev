package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Listen takes a network eg (tcp) and port, returns a listener
	// A listener is an interface that has 3 methods
	// 1. Accept()
	// waits for and returns the next connection to the listener
	// 2. Close()
	// Closes the listener, any blocked Accept operations will be unblocked
	// return errors
	// 3. Addr()
	// Returns listener's network address

	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		// If someone makes a request we accept to get a connection
		// A connection (conn) is an interface,
		// which contains a Read(), Write() ....
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		// Write to a connection
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "how is your day")
		n, err := fmt.Fprintf(conn, "%v", "well, I hope\n")
		fmt.Print(n, " bytes written.\n")
		conn.Close()
	}
}
