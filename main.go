package main

import (
	"fmt"
	"net/http"
)

// http.Handeer

// Anything the implements ServerHTTP(ResponseWriter, *Request)
// is of type handler
/*
	type Handler interface {
		ServerHTTP(ResponseWriter, *Request)
	}

*/

type hotdog int

func main() {
	var d hotdog
	// ListenAndServe(addr string, handler Handler)
	http.ListenAndServe(":8080", d)

}

// Attaching the ServerHTTP(ResponseWriter, *Request) function to
// type hotdog makes hotdog be of type http.Handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this function")
}
