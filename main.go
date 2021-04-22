package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
Package context defines the Context type, which carries deadlines,
cancellation signals, and other request-scoped values across API boundaries
and between processes.

Incoming requests to a server should create a Context, and outgoing calls to
servers should accept a Context. The chain of function calls between them must
propagate the Context, optionally replacing it with a derived Context created
using WithCancel, WithDeadline, WithTimeout, or WithValue.
When a Context is canceled, all Contexts derived from it are also canceled.
*/

/*
When a request comes into a server, you can create a context,
Add some session variable to that context, e.g. session ID, then you can
pass that context around through the system
*/

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	// In practice DO NOT	uses context for storing random values
	// Maybe only session ID or user ID, only values directlry associated with the request
	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	uid, name := dbAccess(ctx)

	fmt.Fprintf(w, "uid: %v \nname: %v \n", uid, name)
}

// Retrieves the Data from the context
// dbAccess returns a tuple
func dbAccess(ctx context.Context) (int, error) {

	// Set to timeout after one second
	// WithTimeout returns a cancel func
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// Create a chanel of type int
	ch := make(chan int)

	go func() {
		// ridiculous long running task
		uid := ctx.Value("userID").(int)
		time.Sleep(10 * time.Second)

		// check to make sure we're not running in vain
		// if ctx.Done() has
		if ctx.Err() != nil {
			return
		}

		ch <- uid
	}()

	select {
	// True if the channel has been cancelled
	case <-ctx.Done():
		return 0, ctx.Err()
	// If the channel produces a value
	case i := <-ch:
		return i, nil
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

// per request variables
// good candidate for putting into context
