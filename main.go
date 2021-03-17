package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// To run program
// 1. $ go run main.go
// 2. $ telnet localhost 8080

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
		go handle(conn)
	}
}

func handle(conn net.Conn) {

	// Set a 10 second deadline no the
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Conn timed out")
	}

	// We use bufio scranner to read from the connection
	scanner := bufio.NewScanner(conn)

	// Scan() returns a bool, it keeps scanning tocken by tocken i.e. scanner contents
	// as we loop until the is nothing to scan the it returns false
	// The token is a line by default
	for scanner.Scan() {
		// Get the text from the tocken
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	// We never get here
	// we have an opent stream connection

	fmt.Println("Code got here")
}
