package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
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
func dbAccess(ctx context.Context) (int, string) {
	// .(int) is an assertion, asserting that ctx.Value("userID") is an int
	uid := ctx.Value("userID").(int)
	name := ctx.Value("fname").(string)
	/*
		acccess DB ...
	*/
	return uid, name
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

// per request variables
// good candidate for putting into context
