package main

import (
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
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
