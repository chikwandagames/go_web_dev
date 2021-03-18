package main

import (
	"io"
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
type hotcat int

func main() {
	var d hotdog
	var c hotcat

	// The trailing / at the end of the path makes a difference
	// "/dog/" will catch  /dog/something/another
	// "/cat" will NOT catch /cat/something
	http.Handle("/dog/", d)
	http.Handle("/cat", c)

	// If you pass nil, ListenAndServe will use DefaultServeMux
	http.ListenAndServe(":8080", nil)
}

// Attaching the ServerHTTP(ResponseWriter, *Request) function to
// type hotdog makes hotdog be of type http.Handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Chihuahua")
}

func (m hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Sphynx")
}
